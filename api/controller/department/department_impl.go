package department

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sectran/api/service"
)

func ListImpl() ([]departmentParameter, int) {
	var tableList []departmentParameter
	var total int
	return tableList, total
}

type DepartmentModel struct {
	ID       string `json:"id"`       //部门ID
	Name     string `json:"name"`     //部门名称
	Describe string `json:"describe"` //部门描述
}

func AddDepartmentImpl(p departmentParameter) error {

	Db := service.Db
	fmt.Printf("%v\n", Db)
	if !Db.HasTable(&DepartmentModel{}) {
		Db.CreateTable(&DepartmentModel{})
	}
	fmt.Printf("%v\n", p.Name)
	var B = DepartmentModel{
		ID:       uuid.Must(uuid.NewV4(), nil).String(),
		Name:     p.Name,
		Describe: p.Describe,
	}
	err := Db.Create(&B).Error

	return err
}
