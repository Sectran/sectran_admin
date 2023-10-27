package auth

import (
	"context"
	"time"

	"sectran/apiservice/internal/middleware"
	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/golang-jwt/jwt"
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
		svcCtx: svcCtx,
	}
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func (l *LoginLogic) Login(req *types.AuthRequest) (resp *types.CommonResponse, err error) {
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": req.Account,
		"nbf":      now,                           // do not work before
		"exp":      now + l.svcCtx.Config.ExpTime, // exp time
		"iat":      now,                           // issur at
		"iss":      "sc@sectran",
	})

	tokenString, err := token.SignedString([]byte(l.svcCtx.Config.Secret))
	if err != nil {
		return types.BuildCommonResponse("null", "erro issue token to this user.", types.ERROR_ISSUE_TOKEN), nil
	}

	l.svcCtx.AuthorizeMiddleware.UserSessionPool[tokenString] = &middleware.UserAuthedInfo{}

	return types.BuildCommonResponse(tokenString, "user login success.", types.REQUEST_SUCCESS), nil
}
