package device

import (
	"context"
	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"sectran_admin/ent/predicate"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
)

func ModifyCheckout(svcCtx *svc.ServiceContext, ctx context.Context, req *types.DeviceInfo) error {
	if req.DepartmentId == nil {
		return types.CustomError("设备所属部门ID不能为空")
	}

	if req.Host == nil {
		return types.CustomError("设备地址不能为空")
	}

	if req.Name == nil {
		return types.CustomError("设备名称不能为空")
	}

	if req.Type == nil {
		return types.CustomError("设备类型不能为空")
	}

	domain := ctx.Value("request_domain").((*ent.User))

	//设备所属部门必须为该主体的子部门
	deviceParentDepartments, err := svcCtx.DB.Department.Query().Where(department.ID(*req.DepartmentId)).Select(department.FieldParentDepartments).String(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return types.CustomError("父部门不存在，可能已被删除")
		}
		return types.ErrInternalError
	}

	//当前主体是否存在权限操作该部门下的设备
	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), deviceParentDepartments); err != nil {
		return err
	}

	//所修改的设备地址同部门中不可以重复
	var predicates []predicate.Department
	if len(deviceParentDepartments) > 0 {
		predicates = append(predicates, department.ParentDepartmentsHasPrefix(deviceParentDepartments))
	} else {
		predicates = append(predicates, department.ParentDepartmentsEQ(deviceParentDepartments))
	}

	//同部门层级不允许出现重复的设备地址，不同部门之间可以
	if c := svcCtx.DB.Device.Query().Where(device.HostEQ(*req.Host)).WithDepartments(func(q *ent.DepartmentQuery) {
		q.Where(predicates...)
	}).CountX(ctx); c > 0 {
		return types.CustomError("当前部门层级存在地址重复的设备")
	}

	return nil
}
