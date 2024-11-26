package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLableTreeByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLableTreeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLableTreeByIdLogic {
	return &GetLableTreeByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *GetLableTreeByIdLogic) GetLableTreeById(req *types.IDReq) (resp *types.LableTreeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
