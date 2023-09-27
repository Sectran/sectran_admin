package department

import (
	"github.com/gin-gonic/gin"
	response "sectran/api"
)

type departmentParameter struct {
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}

// List 查询部门列表
func ListDepartment(c *gin.Context) {
	table, total := ListDepartmentImpl(c)
}

// AddDepartment 添加部门
func AddDepartment(c *gin.Context) {
	p := departmentParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, nil, "请输入")
		return
	}
	err := AddDepartmentImpl(p)
	if err != nil {
		response.RequestOk(c, nil, "添加成功")
	}
}
