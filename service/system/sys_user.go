package system

import (
	"github.com/Sectran/sectran_admin/global"
	"github.com/Sectran/sectran_admin/model/system"
	"github.com/Sectran/sectran_admin/model/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: userInter system.SysUser, err error

type UserService struct{}

// 获取用户列表
func (userService *UserService) ListUser(dto *request.UserListDTO) (total int64, res *[]system.User, err error) {
	limit := dto.Limit
	offset := dto.Offset * (dto.Offset - 1)
	db := global.GVA_DB.Model(&system.User{})

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&res).Error
	if err != nil {
		return
	}

	return
}

// 添加用户
func (userService *UserService) AddUser(user *system.User) (err error) {
	// 当前时间
	return global.GVA_DB.Model(&system.User{}).Create(user).Error
}

// 编辑用户
func (UserService *UserService) UpdateUser(user *system.User) (err error) {
	return global.GVA_DB.Model(user).Where("id=?", user.Id).
		Updates(map[string]interface{}{
			"name":        user.Name,
			"password":    user.Password,
			"dept_id":     user.DeptId,
			"description": user.Description,
			"email":       user.Email,
			"telephone":   user.Telephone,
		}).Error
}

// 删除用户
func (UserService *UserService) DeleteUsers(ids *request.IdDTO) (err error) {
	return global.GVA_DB.Where("id IN ?", ids.Ids).Delete(&system.User{}).Error
}
