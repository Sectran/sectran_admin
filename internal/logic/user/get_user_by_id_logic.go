package user

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"

	dept "sectran_admin/internal/logic/department"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(req *types.IDReq) (*types.UserInfoResp, error) {
	//获取当前主体
	domain := l.ctx.Value("request_domain").((*ent.User))

	data, err := l.svcCtx.DB.User.Get(l.ctx, req.Id)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//查询父部门集合
	uDeptParent, err := l.svcCtx.DB.Department.Query().
		Where(department.ID(uint64(data.DepartmentID))).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err = dept.DomainDeptAccessed((int(domain.DepartmentID)), uDeptParent); err != nil {
		return nil, err
	}

	return &types.UserInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UserInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Account:      &data.Account,
			Name:         &data.Name,
			Password:     &data.Password,
			DepartmentId: &data.DepartmentID,
			RoleId:       &data.RoleID,
			Status:       &data.Status,
			Description:  &data.Description,
			Email:        &data.Email,
			PhoneNumber:  &data.PhoneNumber,
		},
	}, nil
}
