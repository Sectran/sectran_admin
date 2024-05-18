package department

import (
	"bytes"
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
	var (
		err    error
		domain *ent.User
		dDept  *ent.Department
		cDept  *ent.Department
		data   *ent.DepartmentPageList
	)

	defer func(e *error) {
		if *e != nil {
			logx.Errorw("there's an error while get department list", logx.Field("err", *e))
		}
	}(&err)

	domain = l.ctx.Value("request_domain").((*ent.User))
	var predicates []predicate.Department

	//首先查询当前主体的部门、获取到他父亲部门的部门前缀
	dDept, err = l.svcCtx.DB.Department.Get(l.ctx, domain.DepartmentID)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.ErrForceLoginOut
		}
		return nil, types.ErrInternalError
	}

	prefix := fmt.Sprintf("%s%s%d", dDept.ParentDepartments, func() string {
		if dDept.ParentDepartments == "" {
			return ""
		}
		return ","
	}(), dDept.ID)

	if req.ParentDeptId != nil {
		//子集查询必须传递flag
		if req.Flag == nil {
			return nil, types.CustomError("子集查询必须传递Flag字段")
		}

		//判断当前账号是否有权限查询这个部门下的数据
		cDept, err = l.svcCtx.DB.Department.Get(l.ctx, *req.ParentDeptId)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return nil, types.ErrDataNotFound
			}
			return nil, types.ErrInternalError
		}

		//如果当前主体部门的上级部门集合是所请求的部门上级集合的前缀、那么当前账号有权限，否则没有权限
		if !bytes.HasPrefix([]byte(cDept.ParentDepartments), []byte(dDept.ParentDepartments)) {
			return nil, types.ErrAccountHasNoRights
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
	predicates = append(predicates, department.ParentDepartmentsHasPrefix(prefix))

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

	//排序
	deptQuery := l.svcCtx.DB.Department.Query().Where(predicates...)
	data, err = deptQuery.
		Order(department.ByParentDepartments()).
		Order(department.ByID(sql.OrderAsc())).
		Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, types.ErrInternalError
	}

	resp := &types.DepartmentListRespRefer{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	var c int
	HasChildren := func(id uint64) bool {
		//树形结构构建只在一级层级中触发
		if req.ParentDeptId != nil && *req.Flag == 0 {
			c, err = l.svcCtx.DB.Department.Query().Where(department.ParentDepartmentID(id)).Limit(1).Count(l.ctx)
			return (err == nil) && c > 0
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
