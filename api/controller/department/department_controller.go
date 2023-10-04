package department

import (
	"fmt"
	response "sectran/api"
	"sectran/api/common"

	"github.com/gin-gonic/gin"
)

type departmentParameter struct {
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}

// ListDepartment List 查询部门列表
func ListDepartment(c *gin.Context) {
	if err, table, total := listDepartmentImpl(c); err != nil {
		response.RequestError(c, "添加失败")
	} else {
		fmt.Println(table)
		fmt.Println(total)
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
	if err := addDepartmentImpl(p); err != nil {
		response.RequestError(c, "添加失败")
	} else {
		response.RequestOk(c, nil, "添加成功")
	}
}

type RedactDepartmentParameter struct {
	Id string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
	departmentParameter
}

// RedactDepartment 修改部门
func RedactDepartment(c *gin.Context) {
	p := RedactDepartmentParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, "请输入参数")
		return
	}
	if err := redactDepartmentImpl(p); err != nil {
		response.RequestError(c, "修改失败")
	} else {
		response.RequestOk(c, nil, "修改成功")
	}
}

// 删除部门
func DeleteDepartment(c *gin.Context) {
	p := common.DeleteDto{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		return
	}
	_, ok := c.GetPostForm("id")
	if !ok {
		response.RequestError(c, "请输入id")
	}
	if err := deleteDepartmentImpl(p); err != nil {
		response.RequestError(c, "删除失败")
	} else {
		response.RequestOk(c, nil, "删除成功")
	}
}
