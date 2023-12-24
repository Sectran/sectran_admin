package response

import "github.com/Sectran/sectran_admin/model/system"

type TreeDept struct {
	system.Dept
	UserCount   int64
	DeviceCount int64
	Children    []*TreeDept
}
