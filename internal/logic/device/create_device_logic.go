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
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeviceLogic {
	return &CreateDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDeviceLogic) CreateDevice(req *types.DeviceInfo) (*types.BaseMsgResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	if req.DepartmentId == nil {
		return nil, types.CustomError("设备所属部门ID不能为空")
	}

	if req.Host == nil {
		return nil, types.CustomError("设备地址不能为空")
	}

	if req.Name == nil {
		return nil, types.CustomError("设备名称不能为空")
	}

	//设备所属部门必须为该主体的子部门
	deviceParentDepartments, err := l.svcCtx.DB.Department.Query().Where(department.ID(*req.DepartmentId)).Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.CustomError("父部门不存在，可能已被删除")
		}
		return nil, types.ErrInternalError
	}

	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), deviceParentDepartments); err != nil {
		return nil, err
	}

	var predicates []predicate.Department
	if len(deviceParentDepartments) > 0 {
		predicates = append(predicates, department.ParentDepartmentsHasPrefix(deviceParentDepartments))
	} else {
		predicates = append(predicates, department.ParentDepartmentsEQ(deviceParentDepartments))
	}

	//同部门层级不允许出现重复的设备地址，不同部门之间可以
	if c := l.svcCtx.DB.Device.Query().Where(device.HostEQ(*req.Host)).WithDepartments(func(q *ent.DepartmentQuery) {
		q.Where(predicates...)
	}).CountX(l.ctx); c > 0 {
		return nil, types.CustomError("当前部门层级存在地址重复的设备")
	}

	_, err = l.svcCtx.DB.Device.Create().
		SetNotNilName(req.Name).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilHost(req.Host).
		SetNotNilType(req.Type).
		SetNotNilDescription(req.Description).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
