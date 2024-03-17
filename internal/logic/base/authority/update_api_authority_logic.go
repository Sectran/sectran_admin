package authority

import (
	"context"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/zeromicro/go-zero/core/errorx"

	"strconv"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiAuthorityLogic {
	return &UpdateApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiAuthorityLogic) UpdateApiAuthority(req *types.UpdateApiAuthorityReq) (resp *types.BaseMsgResp, err error) {
	role, err := l.svcCtx.DB.Role.Get(l.ctx, req.RoleId)
	if err != nil {
		return nil, err
	}

	roleIDStr := strconv.FormatUint(role.ID, 10)

	// clear old policies
	var oldPolicies [][]string = l.svcCtx.Casbin.GetFilteredPolicy(0, roleIDStr)
	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, roleIDStr)
		if err != nil {
			l.Logger.Errorw("failed to remove roles policy", logx.Field("roleCode", roleIDStr), logx.Field("detail", err.Error()))
			return nil, errorx.NewInvalidArgumentError(err.Error())
		}

		if !removeResult {
			return nil, errorx.NewInvalidArgumentError("casbin.removeFailed")
		}
	}

	// add new policies
	var policies [][]string
	for _, v := range req.Data {
		policies = append(policies, []string{roleIDStr, v.Path, v.Method})
	}

	addResult, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil {
		return nil, errorx.NewInvalidArgumentError("casbin.addFailed")
	}

	if addResult {
		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateFailed)}, nil

}
