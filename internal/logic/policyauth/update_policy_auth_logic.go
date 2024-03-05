package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePolicyAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePolicyAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePolicyAuthLogic {
	return &UpdatePolicyAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePolicyAuthLogic) UpdatePolicyAuth(req *types.PolicyAuthInfo) (*types.BaseMsgResp, error) {
    err := l.svcCtx.DB.PolicyAuth.UpdateOneID(*req.Id).
			SetNotNilName(req.Name).
			SetNotNilPower(req.Power).
			SetNotNilDepartmentID(req.DepartmentId).
			SetNotNilUsers(req.Users).
			SetNotNilAccounts(req.Accounts).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
