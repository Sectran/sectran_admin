package dept

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

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

func (l *DeptEditLogic) DeptEdit(req *types.DeptAllInfo) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
