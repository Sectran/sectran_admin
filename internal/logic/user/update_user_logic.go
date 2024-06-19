package user

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/user"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UserInfo) (*types.BaseMsgResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	targetUser, err := l.svcCtx.DB.User.Query().Where(user.ID(*req.Id)).WithDepartments().Only(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), targetUser.Edges.Departments.ParentDepartments); err != nil {
		return nil, err
	}

	targetDept, err := l.svcCtx.DB.Department.Query().Where(department.ID(*req.DepartmentId)).First(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	if targetDept == nil {
		return nil, types.CustomError("所修改的用户部门不存在")
	}

	// 攻击行为
	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), targetDept.ParentDepartments); err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.User.UpdateOneID(*req.Id).
		SetNotNilAccount(req.Account).
		SetNotNilName(req.Name).
		SetNotNilPassword(req.Password).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilRoleID(req.RoleId).
		SetNotNilStatus(req.Status).
		SetNotNilDescription(req.Description).
		SetNotNilEmail(req.Email).
		SetNotNilPhoneNumber(req.PhoneNumber).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
