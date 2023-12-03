package response

import (
	"github.com/Sectran/sectran_admin/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
