package role

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.RoleAllInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}
	_, err = l.svcCtx.StRoleModel.Insert(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to add this user account", types.ERROR_REUQEST_FAILED), nil
	}

	return types.BuildCommonResponse("null", "user account add successfully", types.REQUEST_SUCCESS), nil
}
