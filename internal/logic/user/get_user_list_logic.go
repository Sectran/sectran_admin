package user

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/predicate"
	"sectran_admin/ent/role"
	"sectran_admin/ent/user"
	deptLogic "sectran_admin/internal/logic/department"
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

func (l *GetUserListLogic) GetUserList(req *types.UserListReqRefer) (*types.UserListRespRefer, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))
	var predicates []predicate.User

	prefix, err := deptLogic.GetCurrentDominDeptPrefix(l.svcCtx, domain)
	if err != nil {
		dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//查询所有子部门下的用户
	predicates = append(predicates, user.HasDepartmentsWith(department.ParentDepartmentsHasPrefix(*prefix)))

	if req.Account != nil {
		predicates = append(predicates, user.AccountContains(*req.Account))
	}

	if req.Name != nil {
		predicates = append(predicates, user.NameContains(*req.Name))
	}

	if req.DepartmentName != nil {
		predicates = append(predicates, user.HasDepartmentsWith(department.NameContains(*req.DepartmentName)))
	}

	if req.RoleName != nil {
		predicates = append(predicates, user.HasRolesWith(role.NameContains(*req.RoleName)))
	}

	if req.Status != nil {
		predicates = append(predicates, user.Status(*req.Status))
	}

	if req.Description != nil {
		predicates = append(predicates, user.DescriptionContains(*req.Description))
	}

	if req.Email != nil {
		predicates = append(predicates, user.EmailContains(*req.Email))
	}

	if req.PhoneNumber != nil {
		predicates = append(predicates, user.PasswordContains(*req.PhoneNumber))
	}

	data, err := l.svcCtx.DB.User.Query().WithRoles().WithDepartments().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UserListRespRefer{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UserInfoRefer{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Account:        &v.Account,
				Name:           &v.Name,
				DepartmentId:   &v.DepartmentID,
				RoleId:         &v.RoleID,
				Status:         &v.Status,
				Description:    &v.Description,
				Email:          &v.Email,
				PhoneNumber:    &v.PhoneNumber,
				RoleName:       &v.Edges.Roles.Name,
				DepartmentName: &v.Edges.Departments.Name,
			})
	}

	return resp, nil
}
