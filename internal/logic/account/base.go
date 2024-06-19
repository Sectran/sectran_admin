package account

import (
	"context"
	"sectran_admin/ent"
	"sectran_admin/ent/account"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	dept "sectran_admin/internal/logic/department"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ProtocolSsh uint8 = 1 + iota
	PtotocolRdp
	PtotocolVnv
	PtotocolSftp
	PtotocolFtp
	PtotocolMysql
	PtotocolOracle
	PtotocolRedis
	ProtocolMax
)

func ModifyCheckout(svcCtx *svc.ServiceContext, ctx context.Context, req *types.AccountInfo) error {
	domain := ctx.Value("request_domain").((*ent.User))

	if req.DeviceId == nil {
		return types.CustomError("设备ID不能为空")
	}
	if req.Username == nil {
		return types.CustomError("账号用户名不能为空")
	}
	if req.Protocol == nil {
		return types.CustomError("账号协议不能为空")
	}
	if *req.Protocol < ProtocolSsh || *req.Protocol > ProtocolMax {
		return types.CustomError("不支持的账号协议")
	}
	if *req.Protocol == ProtocolSsh {
		if req.Password == nil && req.PrivateKey == nil {
			return types.CustomError("账号凭据不能为空（使用密码或者私钥）")
		}
	} else {
		if req.Password == nil {
			return types.CustomError("账号密码不能为空")
		}
	}
	if req.Port == nil {
		return types.CustomError("账号端口不能为空")
	}

	deptId, err := svcCtx.DB.Device.Query().Where(device.ID(*req.DeviceId)).Select(device.FieldDepartmentID).Int(ctx)
	if err != nil {
		logx.Errorw("操作设备账号时查询设备部门失败", logx.Field("DeviceId", *req.DeviceId))
		return types.ErrInternalError
	}

	//设备所属部门必须为该主体的子部门
	deviceParentDepartments, err := svcCtx.DB.Department.Query().Where(department.ID(uint64(deptId))).Select(department.FieldParentDepartments).String(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return types.CustomError("父部门不存在，可能已被删除")
		}
		return types.ErrInternalError
	}

	//当前主体是否存在权限操作该部门下的设备
	if _, err = dept.DomainDeptAccessed(int(domain.DepartmentID), deviceParentDepartments); err != nil {
		return err
	}

	//(三元组：协议、账号、端口)不可重复
	var predicates []predicate.Account
	predicates = append(predicates, account.ProtocolEQ(*req.Protocol))
	predicates = append(predicates, account.UsernameEQ(*req.Username))
	predicates = append(predicates, account.PortEQ(*req.Port))
	predicates = append(predicates, account.DeviceIDEQ(*req.DeviceId))
	acctExt, err := svcCtx.DB.Account.Query().Where(predicates...).Exist(ctx)
	if err != nil {
		logx.Errorw("查询设备账号三元组协议、账号、端口是否重复时失败")
		return types.ErrInternalError
	}

	if acctExt {
		return types.CustomError("该设备下存在重复的三元组：协议、账号、端口")
	}

	return nil
}
