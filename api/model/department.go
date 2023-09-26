package model

type DepartmentModel struct {
	ID       string `json:"id"`       //部门ID
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}
