package device

import (
	"context"

	"sectran_admin/ent/device"
	"sectran_admin/ent/predicate"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetDeviceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDeviceListLogic {
	return &GetDeviceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDeviceListLogic) GetDeviceList(req *types.DeviceListReq) (*types.DeviceListResp, error) {
	var predicates []predicate.Device
	if req.Name != nil {
		predicates = append(predicates, device.NameContains(*req.Name))
	}
	if req.Host != nil {
		predicates = append(predicates, device.HostContains(*req.Host))
	}
	if req.Type != nil {
		predicates = append(predicates, device.TypeContains(*req.Type))
	}
	data, err := l.svcCtx.DB.Device.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.DeviceListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
		types.DeviceInfo{
            BaseIDInfo:    types.BaseIDInfo{
				Id:          &v.ID,
				CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
            },
			Name:	&v.Name,
			DepartmentId:	&v.DepartmentID,
			Host:	&v.Host,
			Type:	&v.Type,
			Description:	&v.Description,
		})
	}

	return resp, nil
}
