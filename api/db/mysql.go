package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sectran/api/common"
)

func init() {

}

func MysqlConnect() {

	var err error
	common.Db, err = gorm.Open("mysql", "root:123456@/db_sectran?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err, 1231)
	}
	//打印查询的sql语句
	common.Db.LogMode(true)
}
