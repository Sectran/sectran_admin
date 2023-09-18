package user

import "sectran/api"

func UserList() string {
	u, err := GetUserById(1)
	if err != nil {
		return api.ResponseError("找不到这个用户")
	}

	s, err := api.ResponseMsg(api.RSP_SUCCECC, "查找用户成功", u)
	if err != nil {
		return api.ResponseError("内部错误")
	}

	return s
}
