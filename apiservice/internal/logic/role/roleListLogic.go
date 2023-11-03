package role

import (
	"context"

	"sectran/apiservice/internal/svc"
	"sectran/apiservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// (*types.CommonResponse, error)
func (l *RoleListLogic) RoleList(req *types.RoleVisibleInfo) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	//err := l.svcCtx.Validator.Struct(req)
	//if err != nil {
	//	return types.BuildCommonResponse("null", "invalid params", types.ERROR_ILLEGAL_PARAMS), nil
	//}
	return
	//roles, err := l.svcCtx.StRoleModel.FindOne(l.ctx, req)
	//if err != nil {
	//	return types.BuildCommonResponse("null", "failed to query users", 501), nil
	//}
	//return types.BuildCommonResponse(roles, "users info query successfully", 200), nil

}
