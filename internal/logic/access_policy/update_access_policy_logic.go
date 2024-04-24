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

type UpdateAccessPolicyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAccessPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAccessPolicyLogic {
	return &UpdateAccessPolicyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAccessPolicyLogic) UpdateAccessPolicy(req *types.AccessPolicyInfo) (*types.BaseMsgResp, error) {
    err := l.svcCtx.DB.AccessPolicy.UpdateOneID(*req.Id).
			SetNotNilName(req.Name).
			SetNotNilPower(req.Power).
			SetNotNilDepartmentID(req.DepartmentId).
			SetNotNilUsers(req.Users).
			SetNotNilAccounts(req.Accounts).
			SetNotNilEffecteTimeStart(pointy.GetTimeMilliPointer(req.EffecteTimeStart)).
			SetNotNilEffecteTimeEnd(pointy.GetTimeMilliPointer(req.EffecteTimeEnd)).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

    return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
