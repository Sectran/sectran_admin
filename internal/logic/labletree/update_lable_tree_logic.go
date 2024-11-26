package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLableTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLableTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLableTreeLogic {
	return &UpdateLableTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *UpdateLableTreeLogic) UpdateLableTree(req *types.LableTreeInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
