package account

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountByIdLogic {
	return &GetAccountByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountByIdLogic) GetAccountById(req *types.IDReqRefer) (*types.AccountInfoResp, error) {
	var (
		err error
	)

	if err = AccountIdCheckout(l.svcCtx, l.ctx, req.Id); err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.Account.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.AccountInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.AccountInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Username: &data.Username,
			Port:     &data.Port,
			Protocol: &data.Protocol,
			DeviceId: &data.DeviceID,
		},
	}

	//认证凭据
	if req.Detail {
		resp.Data.Password = &data.Password
		resp.Data.PrivateKey = &data.PrivateKey
		resp.Data.PrivateKeyPassword = &data.PrivateKeyPassword
	}

	return resp, nil
}
