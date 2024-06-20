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

	prefix := fmt.Sprintf("%s%s%d", pDept.ParentDepartments, func() string {
		if pDept.ParentDepartments == "" {
			return ""
		}
		return ","
	}(), pDept.ID)
	req.ParentDepartments = &prefix

	var sameLevelDeptNames []struct {
		Name string `json:"name"`
	}

	var predicates []predicate.Department
	predicates = append(predicates, department.ParentDepartmentsHasPrefix(prefix))

	//只有编辑才会传递ID
	if req.Id != nil {
		predicates = append(predicates, department.IDNEQ(*req.Id))
	}

	err = svcCtx.DB.Department.Query().
		Where(predicates...).
		Select(department.FieldName).
		Scan(ctx, &sameLevelDeptNames)
	if err != nil {
		return types.ErrInternalError
	}

	//部门名称不能和同层级的部门名称重复
	for _, v := range sameLevelDeptNames {
		if strings.EqualFold(v.Name, *req.Name) {
			return types.CustomError("当前部门层级已经存在相同名称的部门")
		}
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err = DomainDeptAccessed(int(domain.DepartmentID), *req.ParentDepartments); err != nil {
		return err
	}

	return nil

}
