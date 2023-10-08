package department

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
	"sectran/api/common"
)

type departmentParameter struct {
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}

// ListDepartment List 查询部门列表
func ListDepartment(c *gin.Context) {
	if err, table, total := listDepartmentImpl(c); err != nil {
		response.RequestError(c, "查询失败")
	} else {
		data := common.TableDto{Table: table, Total: total}
		response.RequestOk(c, data, "查询成功")
	}
}

// AddDepartment 添加部门
func AddDepartment(c *gin.Context) {
	p := departmentParameter{}

	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if err := addDepartmentImpl(p, c); err != nil {
		response.RequestError(c, "添加失败")
	} else {
		response.RequestOk(c, nil, "添加成功")
	}
}

type EditDepartmentParameter struct {
	Id string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
	departmentParameter
}

// EditDepartment 修改部门
func EditDepartment(c *gin.Context) {
	p := EditDepartmentParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if err := editDepartmentImpl(p); err != nil {
		response.RequestError(c, "修改失败")
	} else {
		response.RequestOk(c, nil, "修改成功")
	}
}

// DeleteDepartment 删除部门
func DeleteDepartment(c *gin.Context) {
	p := common.DeleteDto{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if len(p.Id) == 0 {
		response.RequestError(c, "请输入id")
		return
	}
	if err := deleteDepartmentImpl(p); err != nil {
		response.RequestError(c, "删除失败")
	} else {
		response.RequestOk(c, nil, "删除成功")
	}
}
