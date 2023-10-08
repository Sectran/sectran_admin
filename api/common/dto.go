package common

import "github.com/dgrijalva/jwt-go"

type DeleteDto struct {
	Id string `json:"id" gorm:"type:char(36);primary_key"` //部门ID
}

type TableDto struct {
	Table interface{} `json:"table"`
	Total int         `json:"total"`
}

type TokenDto struct {
	Token string `json:"token"`
}

// JWTClaims token属性
type JWTClaims struct {
	jwt.StandardClaims      // 包中自带的默认属性
	UserID             uint `json:"uid"` // 用户id
}
