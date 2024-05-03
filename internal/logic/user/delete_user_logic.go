package user

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/ent/user"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	dept "sectran_admin/internal/logic/department"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.IDsReq) (*types.BaseMsgResp, error) {
	//获取当前主体
	domain := l.ctx.Value("request_domain").((*ent.User))

	//校验当前主体是否有权限操作待删除的用户
	for _, v := range req.Ids {
		uDeptId, err := l.svcCtx.DB.User.Query().
			Where(user.ID(v)).
			Select(user.FieldDepartmentID).Int(l.ctx)
		if err != nil {
			return nil, types.ErrInternalError
		}

		uDeptParent, err := l.svcCtx.DB.Department.Query().
			Where(department.ID(uint64(uDeptId))).
			Select(department.FieldParentDepartments).String(l.ctx)
		if err != nil {
			return nil, types.ErrInternalError
		}

		//判断当前账号是否对待操作部门存在访问权限
		if _, err = dept.DomainDeptAccessed((int(domain.DepartmentID)), uDeptParent); err != nil {
			return nil, err
		}
	}

	//todo:被删除的用户如果是在线状态应该强制下线
	_, err := l.svcCtx.DB.User.Delete().Where(user.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
