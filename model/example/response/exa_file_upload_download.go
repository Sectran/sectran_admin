package response

import "github.com/Sectran/sectran_admin/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
