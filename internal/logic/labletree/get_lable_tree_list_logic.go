package labletree

import (
	"context"

	"sectran_admin/ent/labletree"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLableTreeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLableTreeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLableTreeListLogic {
	return &GetLableTreeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLableTreeListLogic) GetLableTreeList(req *types.LableTreeListReq) (*types.LableTreeListResp, error) {
	var predicates []predicate.LableTree
	if req.Name != nil {
		predicates = append(predicates, labletree.NameContains(*req.Name))
	}
	if req.Icon != nil {
		predicates = append(predicates, labletree.IconContains(*req.Icon))
	}
	if req.Content != nil {
		predicates = append(predicates, labletree.ContentContains(*req.Content))
	}
	data, err := l.svcCtx.DB.LableTree.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.LableTreeListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.LableTreeInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
            },
			Name:	&v.Name,
			Type:	&v.Type,
			Icon:	&v.Icon,
			Content:	&v.Content,
			ParentLable:	&v.ParentLable,
			LableTargetType:	&v.LableTargetType,
			ParentLables:	&v.ParentLables,
			LableOwner:	&v.LableOwner,
			Inherit:	&v.Inherit,
			RelatedLables:	&v.RelatedLables,
			Description:	&v.Description,
			Ext1:	&v.Ext1,
			Ext2:	&v.Ext2,
		})
	}

	return resp, nil
}
