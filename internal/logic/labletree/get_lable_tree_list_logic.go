package LableTree

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
	if req.ParentLables != nil {
		predicates = append(predicates, labletree.ParentLablesContains(*req.ParentLables))
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
			ParentLable:	&v.ParentLable,
			ParentLables:	&v.ParentLables,
			LableOwner:	&v.LableOwner,
			Inherit:	&v.Inherit,
			RelatedLabels:	&v.RelatedLabels,
			Description:	&v.Description,
		})
	}

	return resp, nil
}
