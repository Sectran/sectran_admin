package dept

import (
	"context"
	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptAddLogic {
	return &DeptAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptAddLogic) DeptAdd(req *types.DeptAddRequest) (*types.CommonResponse, error) {
	err := l.svcCtx.Validator.Struct(req)
	if err != nil {
		return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	}

	_, err = l.svcCtx.StDeptModel.Insert(l.ctx, req)
	if err != nil {
		return types.BuildCommonResponse("null", "failed to add this user account", types.ERROR_REUQEST_FAILED), nil
	}

	return types.BuildCommonResponse("null", "dept account add successfully", types.REQUEST_SUCCESS), nil
}
