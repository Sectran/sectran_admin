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
func List(c *gin.Context) {
	//table, total := ListImpl()
}

func AddDepartment(c *gin.Context) {
	p := departmentParameter{}
	if err := c.ShouldBindJSON(&p); err != nil {
		response.RequestError(c, nil, "请输入")
		return
	}
	AddDepartmentImpl(p)
	//if err != nil {
	//	response.RequestOk(c, nil, "添加成功")
	//}
}
