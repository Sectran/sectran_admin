package response

import "github.com/Sectran/sectran_admin/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId uint                `json:"oldAuthorityId"` // 旧角色ID
}
