package system

import (
	"time"
)

// 部门
type Dept struct {
	Id          int
	Name        string
	Description string
	Seq         string
	CreateTime  time.Time
}

func (Dept) TableName() string {
	return "sys_dept"
}
