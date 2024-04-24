package department

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentByIdLogic {
	return &GetDepartmentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentByIdLogic) GetDepartmentById(req *types.IDReq) (*types.DepartmentInfoResp, error) {
	//查询当前主体的部门、获取到他父亲部门的部门前缀
	domain := l.ctx.Value("request_domain").((*ent.User))
	domainParentDepartments, err := l.svcCtx.DB.Department.Query().
		Where(department.ID(domain.DepartmentID)).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//查询目标的部门
	data, err := l.svcCtx.DB.Department.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err := DomainDeptAccessed(l.ctx, l.svcCtx, domainParentDepartments, data.ParentDepartments); err != nil {
		return nil, err
	}

	return &types.DepartmentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DepartmentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Name:               &data.Name,
			Area:               &data.Area,
			Description:        &data.Description,
			ParentDepartmentId: &data.ParentDepartmentID,
			ParentDepartments:  &data.ParentDepartments,
		},
	}, nil
}
