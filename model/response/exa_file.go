package response

import "innovation/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
