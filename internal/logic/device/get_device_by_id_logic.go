package device

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeviceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeviceByIdLogic {
	return &GetDeviceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeviceByIdLogic) GetDeviceById(req *types.IDReq) (*types.DeviceInfoResp, error) {
	//获取当前主体
	domain := l.ctx.Value("request_domain").((*ent.User))
	data, err := l.svcCtx.DB.Device.Get(l.ctx, req.Id)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//查询父部门集合
	uDeptParent, err := l.svcCtx.DB.Department.Query().
		Where(department.ID(uint64(data.DepartmentID))).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//判断当前账号是否对待操作设备存在访问权限
	if _, err = dept.DomainDeptAccessed((int(domain.DepartmentID)), uDeptParent); err != nil {
		return nil, err
	}

	return &types.DeviceInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DeviceInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Name:         &data.Name,
			DepartmentId: &data.DepartmentID,
			Host:         &data.Host,
			Type:         &data.Type,
			Description:  &data.Description,
		},
	}, nil
}
