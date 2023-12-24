package system

import (
	"github.com/Sectran/sectran_admin/model/common"
	"github.com/Sectran/sectran_admin/model/system"
	"github.com/Sectran/sectran_admin/model/system/request"
	"github.com/gin-gonic/gin"
)

type DeviceApi struct{}

// 查询用户
func (d *DeviceApi) ListDevice(c *gin.Context) {
	var param request.DeviceListDTO
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}
	total, res, err := deviceService.ListDevice(&param)
	if err != nil {
		return
	}
	common.ResponseList(c, res, total)
}

func (d *DeviceApi) AddDevice(c *gin.Context) {
	var param system.Device
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}

	err = deviceService.AddDevice(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (d *DeviceApi) UpdateDevice(c *gin.Context) {
	var param system.Device
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.UpdateDevice(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (d *DeviceApi) DeleteDevice(c *gin.Context) {
	ids := request.IdDTO{}
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.DeleteDevice(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

// 查询用户
func (d *DeviceApi) ListDeviceAccount(c *gin.Context) {
	var param request.DeviceAccountListDTO
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}
	total, res, err := deviceService.ListDeviceAccount(&param)
	if err != nil {
		return
	}
	common.ResponseList(c, res, total)
}

func (d *DeviceApi) AddDeviceAccount(c *gin.Context) {
	var param system.DeviceAccount
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}

	err = deviceService.AddDeviceAccount(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (d *DeviceApi) UpdateDeviceAccount(c *gin.Context) {
	var param system.DeviceAccount
	err := c.ShouldBindJSON(&param)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.UpdateDeviceAccount(&param)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (d *DeviceApi) DeleteDeviceAccount(c *gin.Context) {
	ids := request.IdDTO{}
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = deviceService.DeleteDeviceAccount(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}
