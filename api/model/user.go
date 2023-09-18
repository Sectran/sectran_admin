package model

type User struct {
	Username  string `json:"username"`
	Realname  string `json:"realname"`
	Adreess   string `json:"adreess"`
	Telephone string `json:"telephone"`
	Age       uint8  `json:"age"`
}
