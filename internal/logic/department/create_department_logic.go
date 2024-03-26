package department

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(req *types.DepartmentInfo) (*types.DepartmentInfoResp, error) {
	data, err := l.svcCtx.DB.Department.Create().
		SetNotNilName(req.Name).
		SetNotNilArea(req.Area).
		SetNotNilDescription(req.Description).
		SetNotNilParentDepartmentID(req.ParentDepartmentId).
		SetNotNilParentDepartments(req.ParentDepartments).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
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
