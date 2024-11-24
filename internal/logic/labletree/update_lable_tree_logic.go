package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLableTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLableTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLableTreeLogic {
	return &UpdateLableTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLableTreeLogic) UpdateLableTree(req *types.LableTreeInfo) (*types.BaseMsgResp, error) {
    err := l.svcCtx.DB.LableTree.UpdateOneID(*req.Id).
			SetNotNilName(req.Name).
			SetNotNilType(req.Type).
			SetNotNilIcon(req.Icon).
			SetNotNilContent(req.Content).
			SetNotNilParentLable(req.ParentLable).
			SetNotNilLableTargetType(req.LableTargetType).
			SetNotNilParentLables(req.ParentLables).
			SetNotNilLableOwner(req.LableOwner).
			SetNotNilInherit(req.Inherit).
			SetNotNilRelatedLables(req.RelatedLables).
			SetNotNilDescription(req.Description).
			SetNotNilExt1(req.Ext1).
			SetNotNilExt2(req.Ext2).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
