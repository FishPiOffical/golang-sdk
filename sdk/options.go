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

	// 获取结构体的所有字段定义
	structFields := getStructFields(v)

	// 查找原文中存在但结构体定义中不存在的字段
	for key := range rawData {
		if !containsField(structFields, key) {
			logger.Warn("JSON字段在结构体定义中不存在",
				slog.String("field", key),
				slog.String("type", fmt.Sprintf("%T", v)),
				slog.Any("value", rawData[key]))
		}
	}
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

// getStructFields 获取结构体的所有字段名（包括json tag）
func getStructFields(v interface{}) map[string]bool {
	fields := make(map[string]bool)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fields
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 添加字段名
		fields[field.Name] = true

		// 添加json tag名
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			tagName := strings.Split(jsonTag, ",")[0]
			if tagName != "" && tagName != "-" {
				fields[tagName] = true
			}
		}

		// 如果是嵌套结构体，递归处理
		if field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			fieldVal := val.Field(i)
			if field.Type.Kind() == reflect.Ptr {
				if !fieldVal.IsNil() {
					nestedFields := getStructFields(fieldVal.Interface())
					for k := range nestedFields {
						fields[k] = true
					}
				}
			} else {
				nestedFields := getStructFields(fieldVal.Addr().Interface())
				for k := range nestedFields {
					fields[k] = true
				}
			}
		}
	}

	return fields
}

// containsField 检查字段是否存在于字段集合中
func containsField(fields map[string]bool, key string) bool {
	// 直接匹配
	if fields[key] {
		return true
	}

	// 尝试驼峰和下划线转换
	camelKey := toCamelCase(key)
	if fields[camelKey] {
		return true
	}

	snakeKey := toSnakeCase(key)
	if fields[snakeKey] {
		return true
	}

	return false
}

// toCamelCase 将下划线命名转换为驼峰命名
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// toSnakeCase 将驼峰命名转换为下划线命名
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}
