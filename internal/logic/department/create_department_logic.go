package department

import (
	"context"
	"fmt"
	"strings"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(req *types.DepartmentInfo) (*types.DepartmentInfoResp, error) {
	var (
		err   error
		pDept *ent.Department
		data  *ent.Department
	)

	defer func(e *error) {
		if *e != nil {
			logx.Errorw("there's an error while creating department", logx.Field("err", *e))
		}
	}(&err)

	//查询父部门信息
	pDept, err = l.svcCtx.DB.Department.Get(l.ctx, *req.ParentDepartmentId)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, types.CustomError("父部门不存在，可能已被删除")
		}
		return nil, types.ErrInternalError
	}

	//赋值拼接ParentDepartments
	prefix := fmt.Sprintf("%s%s%d", pDept.ParentDepartments, func() string {
		if pDept.ParentDepartments == "" {
			return ""
		}
		return ","
	}(), pDept.ID)
	req.ParentDepartments = &prefix

	var sameLevelDeptNames []struct {
		Name string `json:"name"`
	}

	err = l.svcCtx.DB.Department.Query().
		Where(department.ParentDepartmentsHasPrefix(prefix)).
		Select(department.FieldName).
		Scan(l.ctx, &sameLevelDeptNames)
	if err != nil {
		return nil, types.ErrInternalError
	}

	//部门名称不能和同层级的部门名称重复
	for _, v := range sameLevelDeptNames {
		if strings.EqualFold(v.Name, *req.Name) {
			return nil, types.CustomError("当前部门层级已经存在相同名称的部门")
		}
	}

	data, err = l.svcCtx.DB.Department.Create().
		SetNotNilName(req.Name).
		SetNotNilArea(req.Area).
		SetNotNilDescription(req.Description).
		SetNotNilParentDepartmentID(req.ParentDepartmentId).
		SetNotNilParentDepartments(req.ParentDepartments).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
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
		},
	}, nil
}
