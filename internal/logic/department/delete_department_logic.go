package department

import (
	"context"
	"fmt"
	"sort"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"
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
	var (
		err               error
		prefix            string
		currParentDeptIds string
	)

	//查询当前主体的部门、获取到他父亲部门的部门前缀
	domain := l.ctx.Value("request_domain").((*ent.User))
	domainParentDepartments, err := l.svcCtx.DB.Department.Query().
		Where(department.ID(domain.DepartmentID)).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//因为创建的id是递增的、我们从后往前删除，并且删除所有当前部门的子部门
	//首先将部门id数组降序排列、即id递减
	sort.Slice(req.Ids, func(i, j int) bool {
		return req.Ids[i] > req.Ids[j]
	})

	//如果最小的id是1、说明是根部门、不允许删除
	if req.Ids[len(req.Ids)-1] == 1 {
		return nil, errorx.NewCodeAbortedError("不允许删除根部门")
	}

	//在事务模块中删除
	if err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		for _, d := range req.Ids {
			currParentDeptIds, err = l.svcCtx.DB.Department.
				Query().
				Where(department.ID(d)).
				Select(department.FieldParentDepartments).String(l.ctx)
			if err != nil {
				//如果目标不存在、跳过
				if _, ok := err.(*ent.NotFoundError); ok {
					continue
				}

				return dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			//校验是否有操作权限
			if _, err = DomainDeptAccessed(l.ctx, l.svcCtx, domainParentDepartments, currParentDeptIds); err != nil {
				return err
			}

			//按照ParentDepartments前缀匹配删除当前部门的所有子部门(会走索引)
			prefix = fmt.Sprintf("%s,%d", currParentDeptIds, d)
			_, err = tx.Department.Delete().Where(department.ParentDepartmentsHasPrefix(prefix)).Exec(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, req)
			}

			//删除当前部门
			_, err = tx.Department.Delete().Where(department.IDEQ(d)).Exec(l.ctx)
			if err != nil {
				return dberrorhandler.DefaultEntError(l.Logger, err, req)
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	//TODO:是否一并删除部门下的各种资源
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
