package sdk

import (
	"fmt"

	"github.com/FishPiOffical/golang-sdk/types"
)

// PostUploadFile 上传文件
func (s *FishPiSDK) PostUploadFile(filePaths []string) (*types.ApiResponse[*types.UploadFileResponse], error) {
	response := new(types.ApiResponse[*types.UploadFileResponse])

	request := s.client.R()

	for _, filePath := range filePaths {
		request.SetFile("file[]", filePath)
	}

	_, err := request.
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/upload")

	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	return response, nil
}
