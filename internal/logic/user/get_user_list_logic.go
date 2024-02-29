package user

import (
	"context"

	"sectran_admin/ent/user"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (*types.UserListResp, error) {
	var predicates []predicate.User
	if req.Account != nil {
		predicates = append(predicates, user.AccountContains(*req.Account))
	}
	if req.Name != nil {
		predicates = append(predicates, user.NameContains(*req.Name))
	}
	if req.Password != nil {
		predicates = append(predicates, user.PasswordContains(*req.Password))
	}
	data, err := l.svcCtx.DB.User.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UserListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.UserInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
            },
			Account:	&v.Account,
			Name:	&v.Name,
			Password:	&v.Password,
			DepartmentId:	&v.DepartmentID,
			RoleId:	&v.RoleID,
			Status:	&v.Status,
			Description:	&v.Description,
			Email:	&v.Email,
			PhoneNumber:	&v.PhoneNumber,
		})
	}

	return resp, nil
}
