package department

import (
	"context"
	"fmt"

	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetChildrenDepartmentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChildrenDepartmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChildrenDepartmentByIdLogic {
	return &GetChildrenDepartmentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChildrenDepartmentByIdLogic) GetChildrenDepartmentById(req *types.ChildrenReq) (*types.DepartmentListResp, error) {
	var prefix string

	dept, err := l.svcCtx.DB.Department.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if dept.ParentDepartments != "" {
		prefix = fmt.Sprintf("%s,%d", dept.ParentDepartments, dept.ID)
	} else {
		prefix = fmt.Sprint(dept.ID)
	}

	deptQuery := l.svcCtx.DB.Department.Query().
		Where(department.ParentDepartmentsHasPrefix(prefix))

	if req.Deep == 1 {
		deptQuery.Where(department.ParentDepartmentID(req.Id))
	}

	data, err := deptQuery.
		Order(department.ByParentDepartments()).
		Order(department.ByID(sql.OrderAsc())).
		Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.DepartmentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	HasChildren := func(id uint64) bool {
		c, err := l.svcCtx.DB.Department.Query().Where(department.ParentDepartmentID(id)).Count(l.ctx)
		return (err == nil) && c > 0
	}

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
				HasChildren:       HasChildren(v.ID),
			})
	}

	return resp, nil
}
