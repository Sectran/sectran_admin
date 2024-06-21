package LableTree

import (
	"context"

    "sectran_admin/ent/labletree"
    "sectran_admin/internal/svc"
    "sectran_admin/internal/types"
    "sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteLableTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLableTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLableTreeLogic {
	return &DeleteLableTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLableTreeLogic) DeleteLableTree(req *types.IDsReq) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.LableTree.Delete().Where(labletree.IDIn(req.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
