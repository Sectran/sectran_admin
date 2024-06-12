package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyAuthByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPolicyAuthByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyAuthByIdLogic {
	return &GetPolicyAuthByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetPolicyAuthByIdLogic) GetPolicyAuthById(req *types.IDReq) (resp *types.PolicyAuthInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
