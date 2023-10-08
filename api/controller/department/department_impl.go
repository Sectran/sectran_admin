package department

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"sectran/api/common"
	"sectran/api/model"
	"strconv"
	"time"
)

// 查询部门列表
func listDepartmentImpl(c *gin.Context) (error, []model.DepartmentModel, int) {
	var tableList []model.DepartmentModel
	var total int
	Db := common.Db
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if !Db.HasTable(&model.DepartmentModel{}) {
		Db.CreateTable(&model.DepartmentModel{})
	}
	//分页
	if page > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	Db = Db.Where("is_delete = 0")
	err := Db.Find(&tableList).Offset(-1).Limit(-1).Count(&total).Error
	return err, tableList, total
}

// 添加部门
func addDepartmentImpl(p departmentParameter, c *gin.Context) error {
	Db := common.Db
	data := common.UserDto{UserName: c.GetString("username")}
	var B = model.DepartmentModel{
		Id:                     uuid.Must(uuid.NewV4(), nil).String(),
		Name:                   p.Name,
		Describe:               p.Describe,
		AddTime:                time.Now().Format("2006-01-02 15:04:05"),
		RevampTime:             time.Now().Format("2006-01-02 15:04:05"),
		AddUser:                data.UserName,
		CorrelationResourceInt: 0,
		CorrelationUserInt:     0,
		IsDelete:               0,
		SuperiorId:             "",
		SubordinateId:          "",
		Location:               "",
	}
	err := Db.Create(&B).Error
	return err
}

// 编辑部门
func editDepartmentImpl(p EditDepartmentParameter) error {
	Db := common.Db
	var tableList []model.DepartmentModel
	fmt.Printf("%v\n", Db)
	err := Db.Model(&tableList).Where("id = ?", p.Id).Updates(map[string]interface{}{"name": p.Name, "describe": p.Describe, "revamp_time": time.Now().Format("2006-01-02 15:04:05")}).Error
	return err
}

// 删除部门
func deleteDepartmentImpl(p common.DeleteDto) error {
	Db := common.Db
	var tableList []model.DepartmentModel
	err := Db.Model(&tableList).Where("id = ?", p.Id).Updates(map[string]interface{}{"is_delete": 1}).Error
	return err
}
