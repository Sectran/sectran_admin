package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyAuthListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyAuthListLogic {
	return &GetPolicyAuthListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPolicyAuthListLogic) GetPolicyAuthList(req *types.PolicyAuthListReq) (resp *types.PolicyAuthListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
