package labletree

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLableTreeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLableTreeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLableTreeByIdLogic {
	return &GetLableTreeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLableTreeByIdLogic) GetLableTreeById(req *types.IDReq) (*types.LableTreeInfoResp, error) {
	data, err := l.svcCtx.DB.LableTree.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.LableTreeInfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
        },
        Data: types.LableTreeInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &data.ID,
				CreatedAt:    pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.UnixMilli()),
            },
			Name:	&data.Name,
			Type:	&data.Type,
			Icon:	&data.Icon,
			Content:	&data.Content,
			ParentLable:	&data.ParentLable,
			LableTargetType:	&data.LableTargetType,
			ParentLables:	&data.ParentLables,
			LableOwner:	&data.LableOwner,
			Inherit:	&data.Inherit,
			RelatedLables:	&data.RelatedLables,
			Description:	&data.Description,
			Ext1:	&data.Ext1,
			Ext2:	&data.Ext2,
        },
	}, nil
}

