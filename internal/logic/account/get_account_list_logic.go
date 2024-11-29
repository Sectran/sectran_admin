package account

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/account"
	"sectran_admin/ent/department"
	"sectran_admin/ent/predicate"
	deptLogic "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAccountListLogic {
	return &GetAccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAccountListLogic) GetAccountList(req *types.AccountListReqRefer) (*types.AccountListResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))
	var predicates []predicate.Account

	prefix, err := deptLogic.GetCurrentDominDeptPrefix(l.svcCtx, domain)
	if err != nil {
		dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	predicates = append(predicates, //查询所有子部门下的设备
		account.HasDepartmentsWith(department.ParentDepartmentsHasPrefix(*prefix)))

	if req.Username != nil {
		predicates = append(predicates, account.UsernameContains(*req.Username))
	}

	if req.Port != nil {
		predicates = append(predicates, account.Port(*req.Port))
	}

	if req.DeviceId != nil {
		predicates = append(predicates, account.DeviceID(*req.DeviceId))
	}

	if req.Protocol != nil {
		predicates = append(predicates, account.Protocol(*req.Protocol))
	}

	data, err := l.svcCtx.DB.Account.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.AccountListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.AccountInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        &v.ID,
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Username: &v.Username,
				Port:     &v.Port,
				Protocol: &v.Protocol,
				DeviceId: &v.DeviceID,
				Password: &v.Password,
			})
	}

	return resp, nil
}
