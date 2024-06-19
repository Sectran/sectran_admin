package user

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/dlclark/regexp2"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func isValidPassword(password string) bool {
	// 密码规则：至少包含一个大写字母、一个小写字母、一个数字，且长度在8到20之间
	regex := regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,20}$`, 0)
	match, _ := regex.MatchString(password)
	return match
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (*types.UserInfoResp, error) {
	var (
		err error
	)

	if err = ModifyCheckout(l.svcCtx, l.ctx, req); err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.User.Create().
		SetNotNilAccount(req.Account).
		SetNotNilName(req.Name).
		SetNotNilPassword(req.Password).
		SetNotNilDepartmentID(req.DepartmentId).
		SetNotNilRoleID(req.RoleId).
		SetNotNilStatus(req.Status).
		SetNotNilDescription(req.Description).
		SetNotNilEmail(req.Email).
		SetNotNilPhoneNumber(req.PhoneNumber).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
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
