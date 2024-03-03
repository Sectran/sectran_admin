package device

import (
	"context"

    "sectran_admin/ent/device"
    "sectran_admin/internal/svc"
    "sectran_admin/internal/types"
    "sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeviceLogic {
	return &DeleteDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeviceLogic) DeleteDevice(req *types.IDsReq) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.Device.Delete().Where(device.IDIn(req.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
