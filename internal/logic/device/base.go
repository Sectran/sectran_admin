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

	"github.com/zeromicro/go-zero/core/logx"
)

func DeviceIdCheckout(svcCtx *svc.ServiceContext, ctx context.Context, deviceId uint64) error {
	return DeviceIdsCheckout(svcCtx, ctx, []uint64{deviceId})
}

func DeviceIdsCheckout(svcCtx *svc.ServiceContext, ctx context.Context, deviceIds []uint64) error {
	domain := ctx.Value("request_domain").((*ent.User))

	deptIds := make([]uint64, 0)
	err := svcCtx.DB.Device.Query().Where(device.IDIn(deviceIds...)).Select(device.FieldDepartmentID).Scan(ctx, &deptIds)
	if err != nil {
		logx.Errorw("操作设备账号时查询设备部门失败", logx.Field("DeviceId", deviceIds))
		return types.ErrInternalError
	}

	deviceParentDepartments := make([]string, 0)
	//设备所属部门必须为该主体的子部门
	err = svcCtx.DB.Department.Query().
		Where(department.IDIn(deptIds...)).
		Select(department.FieldParentDepartments).
		Scan(ctx, &deviceParentDepartments)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return types.CustomError("父部门不存在，可能已被删除")
		}
		return types.ErrInternalError
	}

	//当前主体是否存在权限操作该部门下的设备
	for _, v := range deviceParentDepartments {
		if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), v); err != nil {
			return err
		}
	}
	return nil
}

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

	//编辑设备 验证设备关联部门是否有操作权限 && 验证设备是否有操作权限
	//验证主体能否操作设备
	if req.Id != nil {
		if err := DeviceIdCheckout(svcCtx, ctx, *req.Id); err != nil {
			return err
		}
	}

	//验证主体是否能操作设备部门
	deviceParentDepartments, err := svcCtx.DB.Department.Query().
		Where(department.ID(*req.DepartmentId)).
		Select(department.FieldParentDepartments).String(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return types.CustomError("父部门不存在，可能已被删除")
		}
		return types.ErrInternalError
	}

	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), deviceParentDepartments); err != nil {
		return err
	}

	// 只有编辑才会传递ID
	var predicatesDevice []predicate.Device
	predicatesDevice = append(predicatesDevice, device.HostEQ(*req.Host))
	predicatesDevice = append(predicatesDevice, device.DepartmentIDEQ(*req.DepartmentId))
	if req.Id != nil {
		predicatesDevice = append(predicatesDevice, device.IDNEQ(*req.Id))
	}

	// 同部门层级不允许出现重复的设备地址，不同部门之间可以
	existName, err := svcCtx.DB.Device.Query().Where(predicatesDevice...).Exist(ctx)
	if err != nil {
		return types.ErrInternalError
	}

	if existName {
		return types.CustomError("当前部门层级存在地址重复的设备")
	}

	return nil
}
