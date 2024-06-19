package account

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccountLogic {
	return &UpdateAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAccountLogic) UpdateAccount(req *types.AccountInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.Account.UpdateOneID(*req.Id).
		SetNotNilUsername(req.Username).
		SetNotNilPort(req.Port).
		SetNotNilProtocol(req.Protocol).
		SetNotNilPassword(req.Password).
		SetNotNilPrivateKey(req.PrivateKey).
		SetNotNilDeviceID(req.DeviceId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
