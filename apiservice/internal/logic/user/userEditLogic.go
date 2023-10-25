package user

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserEditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserEditLogic {
	return &UserEditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserEditLogic) UserEdit(req *types.UserAllInfo) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
