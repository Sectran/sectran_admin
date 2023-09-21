package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


func init() {
	var DB *gorm.DB
	var err error
	DB, err = gorm.Open("mysql", "root:123456@/db_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	//打印查询的sql语句
	DB.LogMode(true)
}
