package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePolicyAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePolicyAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePolicyAuthLogic {
	return &UpdatePolicyAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdatePolicyAuthLogic) UpdatePolicyAuth(req *types.PolicyAuthInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
