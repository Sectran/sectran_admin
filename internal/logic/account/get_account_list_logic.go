package account

import (
	"context"

	"sectran_admin/ent/account"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountListLogic {
	return &GetAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountListLogic) GetAccountList(req *types.AccountListReq) (*types.AccountListResp, error) {
	var predicates []predicate.Account
	if req.Username != nil {
		predicates = append(predicates, account.UsernameContains(*req.Username))
	}
	if req.Password != nil {
		predicates = append(predicates, account.PasswordContains(*req.Password))
	}
	if req.PrivateKey != nil {
		predicates = append(predicates, account.PrivateKeyContains(*req.PrivateKey))
	}
	data, err := l.svcCtx.DB.Account.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.AccountListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.AccountInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
            },
			Username:	&v.Username,
			Port:	&v.Port,
			Protocol:	&v.Protocol,
			Password:	&v.Password,
			PrivateKey:	&v.PrivateKey,
			DeviceId:	&v.DeviceID,
		})
	}

	return resp, nil
}
