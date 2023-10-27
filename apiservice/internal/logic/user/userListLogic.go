package user

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.UserQueryInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}

	users, err := l.svcCtx.StUserModel.Find(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to query users", 501), nil
	}

	return types.BuildCommonResponse(users, "users info query successfully", 200), nil
}
