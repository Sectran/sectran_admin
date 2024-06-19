package device

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeviceLogic {
	return &UpdateDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDeviceLogic) UpdateDevice(req *types.DeviceInfo) (*types.BaseMsgResp, error) {
	var (
		err error
	)

	if err = ModifyCheckout(l.svcCtx, l.ctx, req); err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.Device.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilHost(req.Host).
		SetNotNilType(req.Type).
		SetNotNilDescription(req.Description).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
