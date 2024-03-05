package policyauth

import (
	"context"

	"sectran_admin/ent/policyauth"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPolicyAuthListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPolicyAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPolicyAuthListLogic {
	return &GetPolicyAuthListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPolicyAuthListLogic) GetPolicyAuthList(req *types.PolicyAuthListReq) (*types.PolicyAuthListResp, error) {
	var predicates []predicate.PolicyAuth
	if req.Name != nil {
		predicates = append(predicates, policyauth.NameContains(*req.Name))
	}
	if req.Users != nil {
		predicates = append(predicates, policyauth.UsersContains(*req.Users))
	}
	if req.Accounts != nil {
		predicates = append(predicates, policyauth.AccountsContains(*req.Accounts))
	}
	data, err := l.svcCtx.DB.PolicyAuth.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.PolicyAuthListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.PolicyAuthInfo{
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
		})
	}

	return resp, nil
}
