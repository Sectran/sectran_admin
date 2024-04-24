package access_policy

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAccessPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAccessPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAccessPolicyLogic {
	return &CreateAccessPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAccessPolicyLogic) CreateAccessPolicy(req *types.AccessPolicyInfo) (*types.BaseMsgResp, error) {
    _, err := l.svcCtx.DB.AccessPolicy.Create().
			SetNotNilName(req.Name).
			SetNotNilPower(req.Power).
			SetNotNilDepartmentID(req.DepartmentId).
			SetNotNilUsers(req.Users).
			SetNotNilAccounts(req.Accounts).
			SetNotNilEffecteTimeStart(pointy.GetTimeMilliPointer(req.EffecteTimeStart)).
			SetNotNilEffecteTimeEnd(pointy.GetTimeMilliPointer(req.EffecteTimeEnd)).
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
