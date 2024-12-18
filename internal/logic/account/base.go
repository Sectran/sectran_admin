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
	//(三元组：协议、账号、端口)不可重复
	var predicates []predicate.Account

	if req.Id != nil {
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
		return types.ErrInternalError
	}

	if acctExt {
		return types.CustomError("该设备下存在重复的三元组：协议、账号、端口")
	}

	return nil
}
