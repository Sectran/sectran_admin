package model

type DepartmentModel struct {
	Id       string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
	Name     string `json:"name"`                                //部门名称
	Describe string `json:"describe"`                            //部门描述
}
