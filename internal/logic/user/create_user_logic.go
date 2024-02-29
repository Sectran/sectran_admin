package user

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

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

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (*types.BaseMsgResp, error) {
    _, err := l.svcCtx.DB.User.Create().
			SetNotNilAccount(req.Account).
			SetNotNilName(req.Name).
			SetNotNilPassword(req.Password).
			SetNotNilDepartmentID(req.DepartmentId).
			SetNotNilRoleID(req.RoleId).
			SetNotNilStatus(req.Status).
			SetNotNilDescription(req.Description).
			SetNotNilEmail(req.Email).
			SetNotNilPhoneNumber(req.PhoneNumber).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
