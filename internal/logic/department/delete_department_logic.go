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
		dept   *ent.Department
		err    error
		prefix string
	)

	//因为id是递增的、我们从后往前删除，并且删除所有当前部门的子部门
	//首先将部门id数组降序排列
	sort.Slice(req.Ids, func(i, j int) bool {
		return req.Ids[i] > req.Ids[j]
	})

	if req.Ids[len(req.Ids)-1] == 1 {
		return nil, errorx.NewCodeAbortedError("不允许删除根部门")
	}

	for _, d := range req.Ids {
		//依次查询他的父亲集合
		dept, err = l.svcCtx.DB.Department.Get(l.ctx, d)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		//按照ParentDepartments前缀匹配删除当前部门的所有子部门(会走索引)
		prefix = fmt.Sprintf("%s,%d", dept.ParentDepartments, dept.ID)
		_, err = l.svcCtx.DB.Department.Delete().Where(department.ParentDepartmentsHasPrefix(prefix)).Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}

		//删除当前部门
		_, err = l.svcCtx.DB.Department.Delete().Where(department.IDEQ(d)).Exec(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
	}

	//TODO:是否删除部门下的资源、是否删除部门下的策略
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
