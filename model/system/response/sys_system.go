package response

import "github.com/Sectran/sectran_admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
