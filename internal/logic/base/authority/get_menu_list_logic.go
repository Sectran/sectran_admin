package authority

import (
	"context"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"strconv"
	"strings"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListAuthorityLogic {
	return &GetMenuListAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListAuthorityLogic) GetMenuListAuthority(req *types.AuthorityRequestInfo) (resp *types.MenuListInfo, err error) {
	role, err := l.svcCtx.DB.Role.Get(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if req.Type > 2 {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument, "非法的参数")
	}

	roleIDStr := strconv.FormatUint(role.ID, 10)
	var policies [][]string = l.svcCtx.Casbin.GetFilteredPolicy(0, roleIDStr)
	var list []string = make([]string, 0)
	for _, v := range policies {
		name := v[1]
		if req.Type == 2 {
			list = append(list, name)
		}

		if req.Type == 0 && strings.Contains(name, ":") {
			list = append(list, name)
		}

		if req.Type == 1 && strings.Contains(name, "/") {
			list = append(list, name)
		}
	}

	return &types.MenuListInfo{
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		Data: list,
	}, nil
}
