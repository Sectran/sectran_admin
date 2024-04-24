package access_policy

import (
	"context"

	"sectran_admin/ent/accesspolicy"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccessPolicyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccessPolicyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccessPolicyListLogic {
	return &GetAccessPolicyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccessPolicyListLogic) GetAccessPolicyList(req *types.AccessPolicyListReq) (*types.AccessPolicyListResp, error) {
	var predicates []predicate.AccessPolicy
	if req.Name != nil {
		predicates = append(predicates, accesspolicy.NameContains(*req.Name))
	}
	if req.Users != nil {
		predicates = append(predicates, accesspolicy.UsersContains(*req.Users))
	}
	if req.Accounts != nil {
		predicates = append(predicates, accesspolicy.AccountsContains(*req.Accounts))
	}
	data, err := l.svcCtx.DB.AccessPolicy.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.AccessPolicyListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.AccessPolicyInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
            },
			Name:	&v.Name,
			Power:	&v.Power,
			DepartmentId:	&v.DepartmentID,
			Users:	&v.Users,
			Accounts:	&v.Accounts,
			EffecteTimeStart:	pointy.GetUnixMilliPointer(v.EffecteTimeStart.UnixMilli()),
			EffecteTimeEnd:	pointy.GetUnixMilliPointer(v.EffecteTimeEnd.UnixMilli()),
		})
	}

	return resp, nil
}
