package api

import (
	"fmt"
	"sectran/api/model"
	"testing"
)

func TestRsp(t *testing.T) {
	u := model.User{
		Username:  "ryan",
		Realname:  "ryan wilson",
		Adreess:   "nanJing",
		Age:       24,
		Telephone: "121212321312",
	}
	user, _ := ResponseMsg(1234, "this is test message", u)
	fmt.Println(user)

	suc := ResponseSuccess("data success")
	fmt.Println(suc)
}
