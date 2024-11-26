package department

import (
	"context"
	"fmt"
	"sort"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"sectran_admin/ent/user"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/entx"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDepartmentLogic {
	return &DeleteDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteDepartmentLogic) DeleteDepartment(req *types.IDsReq) (*types.BaseMsgResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	//因为创建的id是递增的、我们从后往前删除，并且删除所有当前部门的子部门
	//首先将部门id数组降序排列、即id递减
	sort.Slice(req.Ids, func(i, j int) bool {
		return req.Ids[i] > req.Ids[j]
	})

	//如果最小的id是1、说明是根部门、不允许删除
	if req.Ids[len(req.Ids)-1] == 1 {
		return nil, errorx.NewCodeAbortedError("不允许删除根部门")
	}

	if err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		for _, d := range req.Ids {
			currentDept, err := l.svcCtx.DB.Department.
				Query().
				Where(department.ID(d)).
				Select(department.FieldParentDepartments, department.FieldName, department.FieldParentDepartments).First(l.ctx)
			if err != nil {
				if _, ok := err.(*ent.NotFoundError); ok {
					continue //可能已经被删除了，这里直接跳过也合理
				}

				return types.ErrInternalError
			}

			//校验是否有操作权限
			if _, err = DomainDeptAccessed(int(domain.DepartmentID), currentDept.ParentDepartments); err != nil {
				return err
			}

			//如果当前部门下存在关联用户、资源，不允许删除
			ExistUser, err := l.svcCtx.DB.User.Query().Where(user.DepartmentIDEQ(d)).Exist(l.ctx)
			if err != nil {
				return types.ErrExsitBindResource
			}
			if ExistUser {
				return types.CustomError(fmt.Sprintf("部门【%s】中存在未清理的用户", currentDept.Name))
			}

			ExistDevice, err := l.svcCtx.DB.Device.Query().Where(device.DepartmentIDEQ(d)).Exist(l.ctx)
			if err != nil {
				return types.ErrExsitBindResource
			}
			if ExistDevice {
				return types.CustomError(fmt.Sprintf("部门【%s】中存在未清理的设备", currentDept.Name))
			}

			//按照ParentDepartments前缀匹配删除当前部门的所有子部门(会走索引)
			prefix := fmt.Sprintf("%s,%d", currentDept.ParentDepartments, d)
			_, err = tx.Department.Delete().
				Where(
					department.Or(
						department.ParentDepartmentsHasPrefix(prefix),
						department.IDEQ(d),
					),
				).
				Exec(l.ctx)
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
