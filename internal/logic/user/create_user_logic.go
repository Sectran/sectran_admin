package user

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/role"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/dlclark/regexp2"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	dept "sectran_admin/internal/logic/department"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func isValidPassword(password string) bool {
	// 密码规则：至少包含一个大写字母、一个小写字母、一个数字，且长度在8到20之间
	regex := regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,20}$`, 0)
	match, _ := regex.MatchString(password)
	return match
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (*types.UserInfoResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	var (
		roleExist int
		status    bool = true
	)

	targetDept, err := l.svcCtx.DB.Department.Query().Where(department.ID(*req.DepartmentId)).First(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	if targetDept == nil {
		return nil, types.CustomError("所创建用户的部门不存在")
	}

	// 攻击行为
	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), targetDept.ParentDepartments); err != nil {
		return nil, err
	}

	roleExist, err = l.svcCtx.DB.Role.Query().Where(role.ID(*req.RoleId)).Count(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}
	if roleExist == 0 {
		return nil, types.CustomError("所创建用户的角色不存在")
	}

	if !isValidPassword(*req.Password) {
		return nil, types.CustomError("密码必须至少包含一个大写字母、一个小写字母、一个数字,并且长度在8-20之间")
	}

	data, err := l.svcCtx.DB.User.Create().
		SetNotNilAccount(req.Account).
		SetNotNilName(req.Name).
		SetNotNilPassword(req.Password).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilRoleID(req.RoleId).
		SetNotNilStatus(&status).
		SetNotNilDescription(req.Description).
		SetNotNilEmail(req.Email).
		SetNotNilPhoneNumber(req.PhoneNumber).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UserInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Account:      &data.Account,
			Name:         &data.Name,
			Password:     &data.Password,
			DepartmentId: &data.DepartmentID,
			RoleId:       &data.RoleID,
			Status:       &data.Status,
			Description:  &data.Description,
			Email:        &data.Email,
			PhoneNumber:  &data.PhoneNumber,
		},
	}, nil
}
