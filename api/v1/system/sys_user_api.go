package system

import (
	"github.com/Sectran/sectran_admin/model/common"
	"github.com/Sectran/sectran_admin/model/system"
	"github.com/Sectran/sectran_admin/model/system/request"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// 查询用户
func (u *UserApi) ListUser(c *gin.Context) {
	var user request.UserListDTO
	err := c.ShouldBindJSON(&user)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}
	total, res, err := userService.ListUser(&user)
	if err != nil {
		return
	}
	common.ResponseList(c, res, total)
}

func (b *UserApi) AddUser(c *gin.Context) {
	var user system.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		common.FailWithErr(c, common.ErrInvalidParameter)
		return
	}

	err = userService.AddUser(&user)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (b *UserApi) UpdateUser(c *gin.Context) {
	var user system.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.UpdateUser(&user)
	if err != nil {
		common.FailWithErr(c, common.ErrFailedOperation)
		return
	}
	common.Ok(c)
}

func (b *UserApi) DeleteUser(c *gin.Context) {
	ids := request.IdDTO{}
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.DeleteUsers(&ids)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}
