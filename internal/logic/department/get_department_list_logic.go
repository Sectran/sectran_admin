package department

import (
	"context"

	"sectran_admin/ent/department"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentListLogic {
	return &GetDepartmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentListLogic) GetDepartmentList(req *types.DepartmentListReq) (*types.DepartmentListResp, error) {
	var predicates []predicate.Department
	if req.Name != nil {
		predicates = append(predicates, department.NameContains(*req.Name))
	}
	if req.Area != nil {
		predicates = append(predicates, department.AreaContains(*req.Area))
	}
	if req.Description != nil {
		predicates = append(predicates, department.DescriptionContains(*req.Description))
	}

	data, err := l.svcCtx.DB.Department.Query().Where(predicates...).Order(department.ByParentDepartments()).Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.DepartmentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.DepartmentInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Name:              &v.Name,
				Area:              &v.Area,
				Description:       &v.Description,
				ParentDepartments: &v.ParentDepartments,
			})
	}

	return resp, nil
}
