package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLableTreeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLableTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLableTreeListLogic {
	return &GetLableTreeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetLableTreeListLogic) GetLableTreeList(req *types.LableTreeListReq) (resp *types.LableTreeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
