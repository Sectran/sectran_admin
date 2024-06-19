package user

import (
	"context"

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
	var (
		err error
	)

	if err = ModifyCheckout(l.svcCtx, l.ctx, req); err != nil {
		return nil, err
	}

	//用户账号不允许修改
	ex := l.svcCtx.DB.User.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilRoleID(req.RoleId).
		SetNotNilStatus(req.Status).
		SetNotNilDescription(req.Description).
		SetNotNilEmail(req.Email).
		SetNotNilPhoneNumber(req.PhoneNumber)

	if req.Password != nil && len(*req.Password) > 0 {
		ex.SetNotNilPassword(req.Password)
	}

	err = ex.Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
