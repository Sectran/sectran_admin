package department

import (
	"context"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"strings"
)

// 判断当前账号对目标部门是否有操作权限、因为待操作对象一般都会事先查询出来、
// 所以这里直接将targetParentDepartments传递到参数里面防止重复数据库查询
func DomainDeptAccessed(ctx context.Context, svcCtx *svc.ServiceContext, domainParentDepartments, targetParentDepartments string) (bool, error) {
	//当前主体的父部门前缀是待操作目标的父部门前缀的话、说明当前主体对当前部门目标有操作权限
	if strings.HasPrefix(targetParentDepartments, domainParentDepartments) {
		return true, nil
	}

	return false, types.ErrAccountHasNoRights
}
