package device

import (
	"context"
	"fmt"

	"sectran_admin/ent"
	"sectran_admin/ent/account"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	dept "sectran_admin/internal/logic/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/entx"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDeviceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDeviceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDeviceLogic {
	return &DeleteDeviceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDeviceLogic) DeleteDevice(req *types.IDsReq) (*types.BaseMsgResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		for _, v := range req.Ids {
			result, err := l.svcCtx.DB.Device.Query().
				Where(device.ID(v)).
				WithDepartments(func(q *ent.DepartmentQuery) {
					q.Select(department.FieldParentDepartments) // 查询所属部门的上级部门集合
				}).
				Select(device.FieldName).
				Only(l.ctx)
			if err != nil {
				return err
			}

			if _, err = dept.DomainDeptAccessed((int(domain.DepartmentID)),
				result.Edges.Departments.ParentDepartments); err != nil {
				return err
			}

			existAcc, err := l.svcCtx.DB.Account.Query().Where(account.DepartmentIDEQ(v)).Exist(l.ctx)
			if err != nil {
				return types.ErrExsitBindResource
			}
			if existAcc {
				return types.CustomError(fmt.Sprintf("设备【%s】中存在未清理的账号", result.Name))
			}

			_, err = l.svcCtx.DB.Device.Delete().Where(device.IDIn(req.Ids...)).Exec(l.ctx)
			if err != nil {
				return types.ErrInternalError
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
