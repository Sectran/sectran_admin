package department

import (
	"fmt"
	"sectran_admin/internal/types"
	"strings"
)

// 如果主体部门id在目标的上级部门中，可以确定存在访问权限
func DomainDeptAccessed(domainParentDepartments int, targetParentDepartments string) (bool, error) {
	if len(targetParentDepartments) == 0 {
		return true, nil
	}

	for _, id := range strings.Split(targetParentDepartments, ",") {
		if id == fmt.Sprintf("%d", domainParentDepartments) {
			return true, nil
		}
	}
	return false, types.ErrAccountHasNoRights
}
