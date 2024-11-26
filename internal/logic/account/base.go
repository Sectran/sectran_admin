package account

import (
	"context"
	"sectran_admin/ent/account"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"

	dev "sectran_admin/internal/logic/device"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	ProtocolSsh uint8 = 1 + iota
	ProtocolRdp
	ProtocolVnv
	ProtocolSftp
	ProtocolFtp
	ProtocolMysql
	ProtocolOracle
	ProtocolRedis
	ProtocolMax
)

// protocols is an array that stores the string representation of each protocol.
var ProtocolsMap = [...]string{
	"",
	"SSH",
	"RDP",
	"VNV",
	"SFTP",
	"FTP",
	"MySQL",
	"Oracle",
	"Redis",
}

func AccountIdCheckout(svcCtx *svc.ServiceContext, ctx context.Context, accountId uint64) error {
	return AccountIdsCheckout(svcCtx, ctx, []uint64{accountId})
}

func AccountIdsCheckout(svcCtx *svc.ServiceContext, ctx context.Context, accountIds []uint64) error {
	deviceIds := make([]uint64, 0)
	err := svcCtx.DB.Account.Query().Where(account.IDIn(accountIds...)).Select(account.FieldDeviceID).Scan(ctx, &deviceIds)
	if err != nil {
		logx.Errorw("操作设备账号时查询设备ID失败", logx.Field("accountIds", accountIds))
		return types.ErrInternalError
	}

	return dev.DeviceIdsCheckout(svcCtx, ctx, deviceIds)
}

func ModifyCheckout(svcCtx *svc.ServiceContext, ctx context.Context, req *types.AccountInfo) error {
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

	//(三元组：协议、账号、端口)不可重复
	var predicates []predicate.Account

	if req.Id != nil {
		//校验是否有权限操作该账号
		if err := AccountIdCheckout(svcCtx, ctx, *req.Id); err != nil {
			return err
		}

		//同一设备下不是本身账号、新增没有id、只有编辑才有
		predicates = append(predicates, account.IDNEQ(uint64(*req.Id)))
	}

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
