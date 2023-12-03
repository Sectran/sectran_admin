package request

import (
	"github.com/Sectran/sectran_admin/model/common/request"
	"github.com/Sectran/sectran_admin/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
