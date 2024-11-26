package department

import (
	"context"

	"sectran_admin/ent"
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
	domain := l.ctx.Value("request_domain").((*ent.User))

	//查询目标的部门
	data, err := l.svcCtx.DB.Department.Get(l.ctx, req.Id)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.CustomError("查询的部门不存在")
		}
		return nil, types.ErrInternalError
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err = DomainDeptAccessed((int(domain.DepartmentID)), data.ParentDepartments); err != nil {
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
