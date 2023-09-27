package main

import (
	"sectran/api/common"
	"sectran/api/model"
)

func main() {
	u := model.User{
		Username:  "ryan",
		Realname:  "ryan wilson",
		Adreess:   "nanJing",
		Age:       24,
		Telephone: "121212321312",
	}
	common.OK(u, "操作成功")
	//fmt.Println(user)
}
