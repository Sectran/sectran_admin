package dept

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptListLogic) DeptList(DeptQuery *types.DeptQueryInfo) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(DeptQuery)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}
	roles, err := l.svcCtx.StDeptModel.Find(l.ctx, DeptQuery)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to query roles", 501), nil
	}
	return types.BuildCommonResponse(roles, "roles info query successfully", 200), nil
}
