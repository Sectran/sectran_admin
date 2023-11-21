package dept

import (
	"context"

	"sectran/internal/svc"
	"sectran/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptEditLogic {
	return &DeptEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptEditLogic) DeptEdit(req *types.DeptEditInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", 500), err
	}
	err = l.svcCtx.StDeptModel.Update(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to edit this account", 501), err
	}

	return types.BuildCommonResponse("null", "Role account edit successfully", 200), nil
}
