package department

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(req *types.DepartmentInfo) (*types.BaseMsgResp, error) {
	var (
		err error
	)

	if err = ModifyCheckout(l.svcCtx, l.ctx, req); err != nil {
		return nil, err
	}

	// 不允许修改部门的上级部门
	err = l.svcCtx.DB.Department.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilArea(req.Area).
		SetNotNilDescription(req.Description).
		Exec(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
