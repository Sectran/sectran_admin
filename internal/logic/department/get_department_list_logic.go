package department

import (
	"context"
	"fmt"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"entgo.io/ent/dialect/sql"
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

func (l *GetDepartmentListLogic) GetDepartmentList(req *types.DepartmentListReqRefer) (*types.DepartmentListRespRefer, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))
	var predicates []predicate.Department

	if req.ParentDeptId != nil {
		if req.Flag == nil { //子集查询必须传递flag
			return nil, types.CustomError("子集查询必须传递Flag字段")
		}

		//判断当前账号是否有权限查询这个部门下的数据
		target, err := l.svcCtx.DB.Department.Get(l.ctx, *req.ParentDeptId)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, types.ErrDataNotFound
			}
			return nil, types.ErrInternalError
		}

		//不存在访问权限、但是属于这个部门
		_, err = DomainDeptAccessed(int(domain.DepartmentID),
			fmt.Sprintf("%s,%d", target.ParentDepartments, *req.ParentDeptId))
		if err != nil {
			return nil, err
		}

		switch *req.Flag {
		case 0:
			//查询部门一级子部门
			predicates = append(predicates, department.ParentDepartmentID(*req.ParentDeptId))
		case 1:
			//模糊查询这个部门下的所有部门
		default:
			return nil, types.CustomError("Flag值不合法(0 or 1)")
		}
	}

	//所请求的部门ID必须是当前主体的同级或者子级部门，不允许查询不在当前部门管辖范围内的部门数据
	prefix, err := GetCurrentDominDeptPrefix(l.svcCtx, domain)
	predicates = append(predicates, department.ParentDepartmentsHasPrefix(*prefix))

	//模糊查询部门名称
	if req.Name != nil {
		predicates = append(predicates, department.NameContains(*req.Name))
	}

	//模糊查询部门地区
	if req.Area != nil {
		predicates = append(predicates, department.AreaContains(*req.Area))
	}

	//模糊查询部门描述
	if req.Description != nil {
		predicates = append(predicates, department.DescriptionContains(*req.Description))
	}

	deptQuery := l.svcCtx.DB.Department.Query().Where(predicates...)
	data, err := deptQuery.
		Order(department.ByParentDepartments()).
		Order(department.ByID(sql.OrderAsc())).
		Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, types.ErrInternalError
	}

	resp := &types.DepartmentListRespRefer{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	HasChildren := func(id uint64) bool {
		if req.ParentDeptId != nil && *req.Flag == 0 { //这个字段值只需要一级层级中触发
			c, err := l.svcCtx.DB.Department.Query().Where(department.ParentDepartmentID(id)).Exist(l.ctx)
			return (err == nil) && c
		}

		return false
	}

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.DepartmentInfoRefer{
				DepartmentInfo: types.DepartmentInfo{
					BaseIDInfo: types.BaseIDInfo{
						Id:        &v.ID,
						CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
						UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
					},
					Name:               &v.Name,
					Area:               &v.Area,
					Description:        &v.Description,
					ParentDepartmentId: &v.ParentDepartmentID,
					ParentDepartments:  &v.ParentDepartments,
				},
				HasChildren: HasChildren(v.ID),
			})
	}

	return resp, nil
}
