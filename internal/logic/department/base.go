package department

import (
	"context"
	"fmt"
	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"strings"
)

// 如果主体部门id在目标的上级部门中，可以确定存在访问权限
func DomainDeptAccessed(domainParentDepartments int, targetParentDepartments string) (bool, error) {
	if len(targetParentDepartments) == 0 {
		return true, nil
	}

	for _, id := range strings.Split(targetParentDepartments, ",") {
		if id == fmt.Sprintf("%d", domainParentDepartments) {
			return true, nil
		}
	}
	return false, types.ErrAccountHasNoRights
}

func GetCurrentDominDeptPrefix(svcCtx *svc.ServiceContext, domain *ent.User) (*string, error) {
	domainDept, err := svcCtx.DB.Department.Get(context.Background(), domain.DepartmentID)
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("%s,%d", domainDept.ParentDepartments, domainDept.ID)
	return &prefix, nil
}

func ModifyCheckout(svcCtx *svc.ServiceContext, ctx context.Context, req *types.DepartmentInfo) error {
	if req.Area == nil {
		return types.CustomError("部门归属地不能为空")
	}
	if req.Name == nil {
		return types.CustomError("部门名称不能为空")
	}
	if req.ParentDepartmentId == nil {
		return types.CustomError("部门父级部门ID不能为空")
	}

	//查询当前主体的部门、获取到他父亲部门的部门前缀
	domain := ctx.Value("request_domain").((*ent.User))

	//查询父部门信息
	pDept, err := svcCtx.DB.Department.Get(ctx, *req.ParentDepartmentId)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return types.CustomError("父部门不存在，可能已被删除")
		}
		return types.ErrInternalError
	}

	//判断当前账号是否对待操作 部门存在访问权限
	if _, err = DomainDeptAccessed(int(domain.DepartmentID),
		fmt.Sprintf("%s,%d", pDept.ParentDepartments, *req.ParentDepartmentId)); err != nil {
		return err
	}

	if req.Id != nil {
		reqDept, err := svcCtx.DB.Department.Get(ctx, *req.Id)
		if err != nil {
			if _, ok := err.(*ent.NotFoundError); ok {
				return types.CustomError("父部门不存在，可能已被删除")
			}
			return types.ErrInternalError
		}

		if _, err = DomainDeptAccessed(int(domain.DepartmentID), reqDept.ParentDepartments); err != nil {
			return err
		}
	}

	//同层级部门名称不能重复
	var prefix string
	req.ParentDepartments = &prefix
	if pDept.ParentDepartments == string(pDept.ID) {
		prefix = pDept.ParentDepartments
	} else {
		prefix = fmt.Sprintf("%s,%d", pDept.ParentDepartments, pDept.ID)
	}

	// 判断是否存在同层级相同名称的部门
	var predicates []predicate.Department
	predicates = append(predicates, department.ParentDepartmentsHasPrefix(prefix))

	// 编辑时排除当前自身ID
	if req.Id != nil {
		predicates = append(predicates, department.IDNEQ(*req.Id))
	}

	exists, err := svcCtx.DB.Department.Query().
		Where(
			append(predicates, department.NameEQ(*req.Name))...,
		).
		Exist(ctx)
	if err != nil {
		return types.ErrInternalError
	}

	if exists {
		return types.CustomError("当前部门层级已经存在相同名称的部门")
	}

	return nil

}
