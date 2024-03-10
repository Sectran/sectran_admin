package base

import (
	"context"
	"encoding/json"
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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	var predicates []predicate.User

	if req.Username != nil {
		predicates = append(predicates, sql.FieldEQ("account", req.Username))
	}

	user, err := l.svcCtx.DB.User.Query().Where(predicates...).Only(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	dept, err := l.svcCtx.DB.Department.Get(l.ctx, user.DepartmentID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	role, err := l.svcCtx.DB.Role.Get(l.ctx, user.RoleID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	if user.Password != *req.Password {
		return nil, dberrorhandler.DefaultEntError(l.Logger, fmt.Errorf("密码错误"), req)
	}

	token, err := jwt.GenerateTokenUsingHs256(l.svcCtx.Config.Auth.AccessSecret, time.Hour*time.Duration(l.svcCtx.Config.Auth.AccessExpire), user)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, fmt.Errorf("无法为该用户正确授权"), req)
	}

	userM, err := json.Marshal(user)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	status := l.svcCtx.AuthorityMiddleware.Rds.Set(l.ctx, token, userM, time.Second*5)
	if status.Err() != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, status.Err(), req)
	}

	return &types.LoginRes{
		Base: &types.BaseMsgResp{
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
			Code: 0,
		},
		User:     user,
		Token:    token,
		DeptName: dept.Name,
		RoleName: role.Name,
	}, nil
}
