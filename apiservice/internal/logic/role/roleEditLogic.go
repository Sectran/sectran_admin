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

func (l *RoleEditLogic) RoleEdit(req *types.RoleAllInfo) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
