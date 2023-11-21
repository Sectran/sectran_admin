package user

import (
	"context"

	"sectran/internal/svc"
	"sectran/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAllInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}

	_, err = l.svcCtx.StUserModel.Insert(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to add this user account", types.ERROR_REUQEST_FAILED), nil
	}

	return types.BuildCommonResponse("null", "user account add successfully", types.REQUEST_SUCCESS), nil
}
