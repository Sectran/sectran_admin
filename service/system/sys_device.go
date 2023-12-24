package system

import (
	"github.com/Sectran/sectran_admin/global"
	"github.com/Sectran/sectran_admin/model/system"
	"github.com/Sectran/sectran_admin/model/system/request"
)

type DeviceService struct{}

func (deviceService *DeviceService) ListDevice(dto *request.DeviceListDTO) (total int64, res *[]system.Device, err error) {
	limit := dto.Limit
	offset := dto.Offset * (dto.Offset - 1)
	db := global.GVA_DB.Model(&system.Device{})

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return
	}

	return
}

func (deviceService *DeviceService) AddDevice(param *system.Device) (err error) {
	return global.GVA_DB.Model(&system.Device{}).Create(param).Error
}

func (deviceService *DeviceService) UpdateDevice(param *system.Device) (err error) {
	return global.GVA_DB.Model(param).Where("id=?", param.Id).
		Updates(map[string]interface{}{
			"name":        param.Name,
			"address":     param.Address,
			"os_kind":     param.OsKind,
			"encoding":    param.Encoding,
			"dept_id":     param.DeptId,
			"description": param.Description,
		}).Error
}

func (deviceService *DeviceService) DeleteDevice(ids *request.IdDTO) (err error) {
	return global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&system.Device{}).Error
}

func (deviceService *DeviceService) ListDeviceAccount(dto *request.DeviceAccountListDTO) (total int64, res *[]system.DeviceAccount, err error) {
	limit := dto.Limit
	offset := dto.Offset * (dto.Offset - 1)
	db := global.GVA_DB.Model(&system.DeviceAccount{}).Where("device_id = ?", dto.DeviceId)

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return
	}

	return

}

func (deviceService *DeviceService) AddDeviceAccount(param *system.DeviceAccount) (err error) {
	return global.GVA_DB.Model(&system.DeviceAccount{}).Create(param).Error
}

func (deviceService *DeviceService) UpdateDeviceAccount(param *system.DeviceAccount) (err error) {
	return global.GVA_DB.Model(param).Where("id=?", param.Id).
		Updates(map[string]interface{}{
			"username":             param.Username,
			"password":             param.Password,
			"is_administrator":     param.IsAdministrator,
			"protocol":             param.Protocol,
			"port":                 param.Port,
			"private_key_password": param.PrivateKeyPassword,
			"private_key":          param.PrivateKey,
		}).Error
}

func (deviceService *DeviceService) DeleteDeviceAccount(ids *request.IdDTO) (err error) {
	return global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&system.DeviceAccount{}).Error
}
