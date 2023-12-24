package system

import (
	"github.com/Sectran/sectran_admin/model/system/request"
	"github.com/gin-gonic/gin"
)

type DeptApi struct{}

// 获取部门列表
func (d *DeptApi) ListDept(c *gin.Context) {

}

// 添加部门
func (d *DeptApi) AddDept(c *gin.Context) {

}

// 更新部门
func (d *DeptApi) UpdateDept(c *gin.Context) {

}

// 删除部门
func (d *DeptApi) DeleteDept(c *gin.Context) {
	// 参数绑定
	var ids request.IdDTO
	err := c.BindJSON(&ids)
	if err != nil {
		return
	}

	// deptService.DeleteDept(&ids)

}
