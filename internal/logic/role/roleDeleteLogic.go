package role

import (
	"context"

	"sectran/internal/svc"
	"sectran/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.RoleDeleteRequest) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", 500), err
	}
	err = l.svcCtx.StRoleModel.Delete(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to edit this account", 501), err
	}

	return types.BuildCommonResponse("null", "Role account delete successfully", 200), nil
}
