package system

import (
	"time"
)

type User struct {
	Id          int
	Username    string
	Password    string
	Name        string
	DeptId      int
	IsDisable   int
	Description string
	CreatedAt   time.Time `gorm:"column:created_at;comment:展示值"`
	RoleId      int
	Telephone   string
	Email       string
}

func (User) TableName() string {
	return "sys_user"
}
