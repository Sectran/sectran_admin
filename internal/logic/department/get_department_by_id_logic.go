package department

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

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
	var (
		err                     error
		domain                  *ent.User
		domainParentDepartments string
		data                    *ent.Department
	)

	defer func(e *error) {
		if *e != nil {
			logx.Errorw("there's an error while get department by id", logx.Field("err", *e))
		}
	}(&err)

	//查询当前主体的部门、获取到他父亲部门的部门前缀
	domain = l.ctx.Value("request_domain").((*ent.User))
	domainParentDepartments, err = l.svcCtx.DB.Department.Query().
		Where(department.ID(domain.DepartmentID)).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.ErrForceLoginOut
		}
		return nil, types.ErrInternalError
	}

	//查询目标的部门
	data, err = l.svcCtx.DB.Department.Get(l.ctx, req.Id)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.CustomError("查询的部门不存在")
		}
		return nil, types.ErrInternalError
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err = DomainDeptAccessed(l.ctx, l.svcCtx, domainParentDepartments, data.ParentDepartments); err != nil {
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
