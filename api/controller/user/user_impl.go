package user

import "sectran/api/model"

// 首先定义你需要操作这个表的方法
type UserInterface interface {
	GetUserById(int32) (*model.User, error)
	DelUserById(int32) (*model.User, error)
	EditUserById(int32, *model.User) error
}

// 实现这些方法
func GetUserById(int32) (*model.User, error) {
	return new(model.User), nil
}