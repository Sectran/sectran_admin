package device

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeviceLogic {
	return &CreateDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDeviceLogic) CreateDevice(req *types.DeviceInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Device.Create().
		SetNotNilName(req.Name).
		SetNotNilHost(req.Host).
		SetNotNilDescription(req.Description).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
