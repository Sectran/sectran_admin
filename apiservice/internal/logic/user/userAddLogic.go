package user

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"
	"sectran/apiservice/model/st_user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.UserAllInfo) (resp *types.CommonResponse, err error) {
	l.svcCtx.StUserModel.Insert(l.ctx, &st_user.StUser{})
	return &types.CommonResponse{
		Response: types.Response{
			Code: 200,
			Msg:  "this is user add requet",
		},
		Data: "Success to add user!",
	}, nil
}
