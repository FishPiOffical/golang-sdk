package types

// UploadFileResponse 上传文件响应
type UploadFileResponse struct {
	ErrFiles []string          `json:"errFiles"` // 上传失败的文件列表
	SuccMap  map[string]string `json:"succMap"`  // 成功上传的文件URL映射
}
