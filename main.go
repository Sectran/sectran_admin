package main

import (
	"fmt"
	"sectran/api"
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
	user, _ := api.ResponseMsg(1234, "this is test message", u)
	fmt.Println(user)

	suc, _ := api.ResponseSuccess("data success")
	fmt.Println(suc)
}
