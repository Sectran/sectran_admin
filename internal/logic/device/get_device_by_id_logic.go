package device

import (
	"context"

	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeviceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeviceByIdLogic {
	return &GetDeviceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeviceByIdLogic) GetDeviceById(req *types.IDReq) (*types.DeviceInfoResp, error) {
	data, err := l.svcCtx.DB.Device.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.DeviceInfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
        },
        Data: types.DeviceInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &data.ID,
				CreatedAt:    pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(data.UpdatedAt.UnixMilli()),
            },
			Name:	&data.Name,
			DepartmentId:	&data.DepartmentID,
			Host:	&data.Host,
			Type:	&data.Type,
			Description:	&data.Description,
        },
	}, nil
}

