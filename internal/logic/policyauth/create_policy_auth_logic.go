package policyauth

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePolicyAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePolicyAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePolicyAuthLogic {
	return &CreatePolicyAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePolicyAuthLogic) CreatePolicyAuth(req *types.PolicyAuthInfo) (*types.BaseMsgResp, error) {
    _, err := l.svcCtx.DB.PolicyAuth.Create().
			SetNotNilName(req.Name).
			SetNotNilPower(req.Power).
			SetNotNilDepartmentID(req.DepartmentId).
			SetNotNilUsers(req.Users).
			SetNotNilAccounts(req.Accounts).
			SetNotNilDirection(req.Direction).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
