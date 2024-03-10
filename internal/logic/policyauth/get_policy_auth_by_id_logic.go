package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyAuthByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPolicyAuthByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyAuthByIdLogic {
	return &GetPolicyAuthByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPolicyAuthByIdLogic) GetPolicyAuthById(req *types.IDReq) (*types.PolicyAuthInfoResp, error) {
	data, err := l.svcCtx.DB.PolicyAuth.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.PolicyAuthInfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
        },
        Data: types.PolicyAuthInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &data.ID,
				CreatedAt:    pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.UnixMilli()),
            },
			Name:	&data.Name,
			Power:	&data.Power,
			DepartmentId:	&data.DepartmentID,
			Users:	&data.Users,
			Accounts:	&data.Accounts,
			Direction:	&data.Direction,
        },
	}, nil
}

