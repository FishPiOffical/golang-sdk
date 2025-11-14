package sdk

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/imroc/req/v3"
)

// Option SDK选项函数
type Option func(sdk *FishPiSDK)

// WithLogDir 设置日志目录，将请求日志保存到指定目录
func WithLogDir(logDir string) Option {
	return func(sdk *FishPiSDK) {
		sdk.logDir = logDir
		// 确保目录存在
		if err := os.MkdirAll(logDir, 0755); err != nil {
			slog.Error("failed to create log directory", slog.String("dir", logDir), slog.Any("error", err))
			return
		}

		sdk.client.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			// 生成日志文件名: 方法+路径+日期.txt
			method := req.Method
			path := strings.ReplaceAll(req.RawURL, "/", "_")
			if path == "" {
				path = "root"
			}
			date := time.Now().Format("20060102_150405")
			filename := fmt.Sprintf("%s%s_%s.txt", method, path, date)
			logPath := filepath.Join(sdk.logDir, filename)

			// 使用 SetDumpOptions 而不是 SetOutputFile，避免干扰响应读取
			req.EnableDumpToFile(logPath)
			return nil
		})
	}
}

// WithCustomUnmarshaler 设置自定义JSON反序列化方法，会检测字段差异并记录日志
// 处理流程：
// 1. 检测原始JSON数据和结构体定义之间的字段差异
// 2. 执行反序列化
// 3. 检测原始JSON数据和反序列化后的数据之间的差异
func WithCustomUnmarshaler(logger *slog.Logger) Option {
	return func(sdk *FishPiSDK) {
		if logger == nil {
			logger = slog.Default()
		}
		sdk.logger = logger

		sdk.client.SetJsonUnmarshal(func(data []byte, v interface{}) error {
			// 第一步：检测原始JSON和结构体定义之间的字段差异
			logger.Debug("开始检测JSON和结构体定义的字段差异")
			detectStructDefinitionDifferences(data, v, logger)

			// 第二步：执行标准反序列化
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}

			// 第三步：检测原始JSON和反序列化后数据的差异
			logger.Debug("开始检测JSON原文和反序列化后数据的差异")
			detectUnmarshalDifferences(data, v, logger)

			return nil
		})
	}
}

// detectStructDefinitionDifferences 检测JSON原文和结构体定义之间的字段差异
func detectStructDefinitionDifferences(data []byte, v interface{}, logger *slog.Logger) {
	// 解析原始JSON为map
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		// 如果不是对象类型，跳过检测
		return
	}

	// 获取结构体的所有字段定义及其类型信息
	structFieldsMap := getStructFieldsWithType(v)

	// 查找原文中存在但结构体定义中不存在的字段，以及类型不匹配的字段
	checkFieldDifferences(rawData, structFieldsMap, "", logger)
}

// checkFieldDifferences 递归检查字段差异和类型匹配
func checkFieldDifferences(rawData map[string]interface{}, structFields map[string]reflect.Type, path string, logger *slog.Logger) {
	for key, value := range rawData {
		currentPath := key
		if path != "" {
			currentPath = path + "." + key
		}

		fieldType, exists := structFields[key]
		if !exists {
			logger.Warn("JSON字段在结构体定义中不存在",
				slog.String("path", currentPath),
				slog.Any("value", value),
				slog.String("jsonType", fmt.Sprintf("%T", value)))
			continue
		}

		// 检查类型是否匹配
		if !isTypeCompatible(value, fieldType, logger, currentPath) {
			logger.Warn("JSON字段类型与结构体定义不匹配",
				slog.String("path", currentPath),
				slog.String("jsonType", fmt.Sprintf("%T", value)),
				slog.String("structType", fieldType.String()),
				slog.Any("value", value))
		}

		// 如果是嵌套对象，递归检查
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if fieldType.Kind() == reflect.Struct || (fieldType.Kind() == reflect.Ptr && fieldType.Elem().Kind() == reflect.Struct) {
				actualType := fieldType
				if fieldType.Kind() == reflect.Ptr {
					actualType = fieldType.Elem()
				}
				nestedFields := getStructFieldsWithTypeFromType(actualType)
				checkFieldDifferences(nestedMap, nestedFields, currentPath, logger)
			}
		}

		// 如果是数组，检查数组元素
		if arrayVal, ok := value.([]interface{}); ok {
			if fieldType.Kind() == reflect.Slice || fieldType.Kind() == reflect.Array {
				elemType := fieldType.Elem()
				if len(arrayVal) > 0 {
					if elemMap, ok := arrayVal[0].(map[string]interface{}); ok {
						if elemType.Kind() == reflect.Struct || (elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct) {
							actualElemType := elemType
							if elemType.Kind() == reflect.Ptr {
								actualElemType = elemType.Elem()
							}
							elemFields := getStructFieldsWithTypeFromType(actualElemType)
							checkFieldDifferences(elemMap, elemFields, fmt.Sprintf("%s[0]", currentPath), logger)
						}
					}
				}
			}
		}
	}
}

// isTypeCompatible 检查JSON值类型是否与结构体字段类型兼容
func isTypeCompatible(jsonValue interface{}, structType reflect.Type, logger *slog.Logger, path string) bool {
	// 处理指针类型
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	switch v := jsonValue.(type) {
	case bool:
		return structType.Kind() == reflect.Bool
	case float64:
		// JSON中的数字都是float64，可以兼容int/float等数值类型
		return structType.Kind() >= reflect.Int && structType.Kind() <= reflect.Float64
	case string:
		return structType.Kind() == reflect.String
	case map[string]interface{}:
		return structType.Kind() == reflect.Struct || structType.Kind() == reflect.Map
	case []interface{}:
		return structType.Kind() == reflect.Slice || structType.Kind() == reflect.Array
	case nil:
		// nil可以赋值给指针、接口、切片、map等
		return true
	default:
		logger.Debug("未知的JSON类型",
			slog.String("path", path),
			slog.String("type", fmt.Sprintf("%T", v)))
		return true
	}
}

// getStructFieldsWithType 获取结构体的所有字段名和类型信息
func getStructFieldsWithType(v interface{}) map[string]reflect.Type {
	fields := make(map[string]reflect.Type)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fields
	}

	typ := val.Type()
	return getStructFieldsWithTypeFromType(typ)
}

// getStructFieldsWithTypeFromType 从reflect.Type获取字段名和类型信息
func getStructFieldsWithTypeFromType(typ reflect.Type) map[string]reflect.Type {
	fields := make(map[string]reflect.Type)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return fields
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 跳过未导出的字段
		if !field.IsExported() {
			continue
		}

		// 获取json tag作为字段名
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			tagName := strings.Split(jsonTag, ",")[0]
			if tagName != "-" && tagName != "" {
				fields[tagName] = field.Type
			}
		} else {
			// 没有json tag，使用字段名
			fields[field.Name] = field.Type
		}

		// 处理嵌入结构体（匿名字段）
		if field.Anonymous && (field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct)) {
			embeddedType := field.Type
			if embeddedType.Kind() == reflect.Ptr {
				embeddedType = embeddedType.Elem()
			}
			embeddedFields := getStructFieldsWithTypeFromType(embeddedType)
			for k, t := range embeddedFields {
				if _, exists := fields[k]; !exists {
					fields[k] = t
				}
			}
		}
	}

	return fields
}

// detectUnmarshalDifferences 检测原始JSON和反序列化后数据的差异
func detectUnmarshalDifferences(data []byte, v interface{}, logger *slog.Logger) {
	// 解析原始JSON为map
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		return
	}

	// 将反序列化后的结构体转换回JSON再解析为map
	marshaledData, err := json.Marshal(v)
	if err != nil {
		logger.Error("无法序列化反序列化后的数据", slog.Any("error", err))
		return
	}

	var unmarshaledData map[string]interface{}
	if err := json.Unmarshal(marshaledData, &unmarshaledData); err != nil {
		logger.Error("无法解析序列化后的数据", slog.Any("error", err))
		return
	}

	// 比较原始数据和反序列化后的数据
	compareMapDifferences(rawData, unmarshaledData, "", logger)
}

// compareMapDifferences 递归比较两个map的差异
func compareMapDifferences(original, unmarshaled map[string]interface{}, path string, logger *slog.Logger) {
	// 检查原始数据中存在但反序列化后丢失的字段
	for key, originalValue := range original {
		currentPath := key
		if path != "" {
			currentPath = path + "." + key
		}

		unmarshaledValue, exists := unmarshaled[key]
		if !exists {
			logger.Warn("字段在反序列化后丢失",
				slog.String("path", currentPath),
				slog.Any("originalValue", originalValue))
			continue
		}

		// 如果值都是map，递归比较
		if origMap, ok := originalValue.(map[string]interface{}); ok {
			if unmMap, ok := unmarshaledValue.(map[string]interface{}); ok {
				compareMapDifferences(origMap, unmMap, currentPath, logger)
				continue
			}
		}

		// 如果值都是数组，比较长度和元素
		if origArray, ok := originalValue.([]interface{}); ok {
			if unmArray, ok := unmarshaledValue.([]interface{}); ok {
				if len(origArray) != len(unmArray) {
					logger.Warn("数组长度不一致",
						slog.String("path", currentPath),
						slog.Int("originalLength", len(origArray)),
						slog.Int("unmarshaledLength", len(unmArray)))
				}
				// 可以进一步比较数组元素
				for i := 0; i < len(origArray) && i < len(unmArray); i++ {
					if origElem, ok := origArray[i].(map[string]interface{}); ok {
						if unmElem, ok := unmArray[i].(map[string]interface{}); ok {
							compareMapDifferences(origElem, unmElem, fmt.Sprintf("%s[%d]", currentPath, i), logger)
						}
					}
				}
				continue
			}
		}

		// 检查值类型是否改变
		if fmt.Sprintf("%T", originalValue) != fmt.Sprintf("%T", unmarshaledValue) {
			logger.Warn("字段类型在反序列化后改变",
				slog.String("path", currentPath),
				slog.String("originalType", fmt.Sprintf("%T", originalValue)),
				slog.String("unmarshaledType", fmt.Sprintf("%T", unmarshaledValue)),
				slog.Any("originalValue", originalValue),
				slog.Any("unmarshaledValue", unmarshaledValue))
		}
	}

	// 检查反序列化后新增的字段（通常是默认值）
	for key := range unmarshaled {
		currentPath := key
		if path != "" {
			currentPath = path + "." + key
		}

		if _, exists := original[key]; !exists {
			logger.Info("反序列化后新增字段（可能是默认值）",
				slog.String("path", currentPath),
				slog.Any("value", unmarshaled[key]))
		}
	}
}
