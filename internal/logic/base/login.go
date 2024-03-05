package base

import (
	"context"
	"fmt"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"
	"sectran_admin/internal/utils/jwt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *LoginLogic) Login(req *types.LoginInfo) (resp *types.BaseDataInfo, err error) {
	var predicates []predicate.User

	if req.Username != nil {
		predicates = append(predicates, sql.FieldEQ("username", req.Username))
	}

	user, err := l.svcCtx.DB.User.Query().Where(predicates...).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if user.Password != *req.Password {
		return nil, dberrorhandler.DefaultEntError(l.Logger, fmt.Errorf("密码错误"), req)
	}

	token, err := jwt.GenerateTokenUsingHs256(l.svcCtx.Config.Auth.AccessSecret, time.Hour*time.Duration(l.svcCtx.Config.Auth.AccessExpire), user)
	if err == nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, fmt.Errorf("无法为该用户正确授权"), req)
	}

	status := l.svcCtx.Authority.Rds.Set(l.ctx, token, user, l.svcCtx.Config.Signature.Expiry)
	if status.Err() != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, status.Err(), req)
	}

	return &types.BaseDataInfo{
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		Code: 0,
		Data: token,
	}, nil
}
