package UserController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"sectran/api/common"
	"sectran/api/model"
	"strconv"
	"time"
)

// 查询用户列表
func listUserImpl(c *gin.Context) (error, []model.UserModel, int) {
	var tableList []model.UserModel
	var total int
	Db := common.Db
	if !Db.HasTable(&model.UserModel{}) {
		Db.CreateTable(&model.UserModel{})
	}

	Db = Db.Where("is_delete = 0")
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	//分页
	if page > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	err := Db.Find(&tableList).Offset(-1).Limit(-1).Count(&total).Error
	return err, tableList, total
}

// 查询添加部门
func addUserImpl(p common.UserDto) (error, string) {
	Db := common.Db
	if err := Db.Where("user_name = ?", p.UserName).First(&model.UserModel{}).Error; err != nil {
		//密码加密保存到数据库
		hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			fmt.Println(err)
		}
		var b = model.UserModel{
			Id:         uuid.Must(uuid.NewV4(), nil).String(),
			UserName:   p.UserName,
			Password:   string(hash),
			RevampTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		if err := Db.Create(&b).Error; err != nil {
			return err, "账号新建失败"
		} else {
			return err, "账号新建成功"
		}

	} else {
		return err, "已有账号，请勿重复添加"
	}

}

// 编辑用户
func editUserImpl(p EditDepartmentParameter) error {
	Db := common.Db
	var tableList []model.UserModel
	hash, _ := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost) //加密处理
	err := Db.Model(&tableList).Where("id = ?", p.Id).Updates(map[string]interface{}{"user_name": p.UserName, "password": hash, "revamp_time": time.Now().Format("2006-01-02 15:04:05")}).Error
	return err
}

// 删除用户
func deleteUserImpl(p common.DeleteDto) error {
	Db := common.Db
	var tableList []model.UserModel
	err := Db.Model(&tableList).Where("id = ?", p.Id).Updates(map[string]interface{}{"is_delete": 1}).Error
	return err
}
