package role

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(roleQuery *types.RoleQueryInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(roleQuery)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}
	roles, err := l.svcCtx.StRoleModel.Find(l.ctx, roleQuery)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to query users", 501), nil
	}
	return types.BuildCommonResponse(roles, "users info query successfully", 200), nil
}
