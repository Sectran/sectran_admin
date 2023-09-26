package department

import "github.com/gin-gonic/gin"

type ListParameter struct {
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}

// List 查询部门列表
func List(c *gin.Context) {
	table, total := ListImpl()
}
