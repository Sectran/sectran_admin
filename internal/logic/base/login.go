package base

import (
	"context"
	"encoding/json"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
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
		return nil, types.CustomError("用户不存在")
	}

	dept, err := l.svcCtx.DB.Department.Get(l.ctx, user.DepartmentID)
	if err != nil {
		return nil, types.ErrInternalError
	}

	role, err := l.svcCtx.DB.Role.Get(l.ctx, user.RoleID)
	if err != nil {
		return nil, types.ErrInternalError
	}

	if user.Password != *req.Password {
		return nil, types.CustomError("用户认证失败")
	}

	exp := time.Hour * time.Duration(l.svcCtx.Config.Auth.AccessExpire)
	token, err := jwt.GenerateTokenUsingHs256(l.svcCtx.Config.Auth.AccessSecret, exp, user)
	if err != nil {
		return nil, types.CustomError("无法为该用户正确授权,请联系开发者")
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//set timeout 5 secends
	status := l.svcCtx.AuthorityMiddleware.Rds.Set(context.Background(), token, userJson, exp)
	if status.Err() != nil {
		return nil, types.CustomError("系统服务繁忙,请稍后再试")
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
