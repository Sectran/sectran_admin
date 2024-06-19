package device

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeviceLogic {
	return &DeleteDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeviceLogic) DeleteDevice(req *types.IDsReq) (*types.BaseMsgResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	//校验当前主体是否有权限操作待删除的设备资源
	for _, v := range req.Ids {
		//查询当前设备的部门id
		dDeptId, err := l.svcCtx.DB.Device.Query().
			Where(device.ID(v)).
			Select(device.FieldDepartmentID).Int(l.ctx)
		if err != nil {
			return nil, types.ErrInternalError
		}

		//查询当前设备所属部门的上级部门集合
		dDeptParent, err := l.svcCtx.DB.Department.Query().
			Where(department.ID(uint64(dDeptId))).
			Select(department.FieldParentDepartments).String(l.ctx)
		if err != nil {
			return nil, types.ErrInternalError
		}

		//判断当前账号是否对待操作部门存在访问权限
		if _, err = dept.DomainDeptAccessed((int(domain.DepartmentID)), dDeptParent); err != nil {
			return nil, err
		}
	}

	_, err := l.svcCtx.DB.Device.Delete().Where(device.IDIn(req.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//删除设备关联的账号
	//策略中删除设备

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
