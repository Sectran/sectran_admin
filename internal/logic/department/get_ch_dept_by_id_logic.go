package department

import (
	"context"
	"fmt"
	"strings"

	"sectran_admin/ent"
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

	data, err := l.svcCtx.DB.Department.Query().
		Where(department.ParentDepartmentsHasPrefix(prefix)).
		Order(department.ByParentDepartments()).
		Order(department.ByID(sql.OrderAsc())).
		Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.DepartmentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	//des当前遍历结构
	//from遍历开始索引
	//depID当前部门id
	//deep当前部门树深
	hasCh := func(des []*ent.Department, from int, deptID int, deep int) bool {
		var deps string
		for i := from; i < len(des); i++ {
			deps = des[i].ParentDepartments
			//如果当前🌲深已经跨过可能的子部门，那么直接返回
			if (len(deps)+1)/2 > deep+1 {
				return false
			}

			if strings.HasSuffix(des[i].ParentDepartments, fmt.Sprint(deptID)) {
				return true
			}
		}

		return false
	}

	for i, v := range data.List {
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
				HasChildren:       hasCh(data.List, i+1, int(v.ID), (len(v.ParentDepartments)+1)/2),
			})
	}

	return resp, nil
}
