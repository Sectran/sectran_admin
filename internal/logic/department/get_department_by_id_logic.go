package department

import (
	"bytes"
	"context"

	"sectran_admin/ent"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDepartmentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDepartmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDepartmentByIdLogic {
	return &GetDepartmentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDepartmentByIdLogic) GetDepartmentById(req *types.IDReq) (*types.DepartmentInfoResp, error) {
	domain := l.ctx.Value("request_domain").((*ent.User))

	data, err := l.svcCtx.DB.Department.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//首先查询当前主体的部门、获取到他父亲部门的部门前缀
	dDept, err := l.svcCtx.DB.Department.Get(l.ctx, domain.DepartmentID)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//如果当前主体部门的上级部门集合是所请求的部门上级集合的前缀、那么当前账号有权限，否则没有权限
	if !bytes.HasPrefix([]byte(data.ParentDepartments), []byte(dDept.ParentDepartments)) {
		return nil, types.ErrAccountHasNoRights
	}

	return &types.DepartmentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.DepartmentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Name:               &data.Name,
			Area:               &data.Area,
			Description:        &data.Description,
			ParentDepartmentId: &data.ParentDepartmentID,
			ParentDepartments:  &data.ParentDepartments,
		},
	}, nil
}
