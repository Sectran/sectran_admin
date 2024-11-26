package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLableTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLableTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLableTreeLogic {
	return &CreateLableTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *CreateLableTreeLogic) CreateLableTree(req *types.LableTreeInfo) (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
