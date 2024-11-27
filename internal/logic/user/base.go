package user

import (
	"context"
	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/role"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

func ModifyCheckout(svcCtx *svc.ServiceContext, ctx context.Context, req *types.UserInfo) error {
	domain := ctx.Value("request_domain").((*ent.User))
	targetDept, err := svcCtx.DB.Department.Query().Where(department.ID(*req.DepartmentId)).First(ctx)
	if err != nil {
		logx.Errorw("操作账号时查询部门失败", logx.Field("DepartmentId", *req.DepartmentId))
		return types.ErrInternalError
	}

	if targetDept == nil {
		return types.CustomError("操作的账号部门不存在")
	}

	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), targetDept.ParentDepartments); err != nil {
		return err
	}

	roleExt, err := svcCtx.DB.Role.Query().Where(role.ID(*req.RoleId)).Exist(ctx)
	if err != nil {
		logx.Errorw("操作账号时查询角色失败", logx.Field("RoleId", *req.RoleId))
		return types.ErrInternalError
	}

	if !roleExt {
		return types.CustomError("操作的账号角色不存在")
	}

	if req.Password != nil && len(*req.Password) > 0 {
		if !isValidPassword(*req.Password) {
			return types.CustomError("密码必须至少包含一个大写字母、一个小写字母、一个数字,并且长度在8-20之间")
		}
	}

	return nil
}
