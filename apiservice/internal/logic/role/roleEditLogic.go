package role

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleEditLogic {
	return &RoleEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleEditLogic) RoleEdit(req *types.RoleEditInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", 500), err
	}
	err = l.svcCtx.StRoleModel.Update(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to edit this account", 501), err
	}

	return types.BuildCommonResponse("null", "Role account edit successfully", 200), nil
}
