package jwt

import (
	"errors"
	"math/rand"
	"sectran_admin/ent"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	UserID   int
	Username string
	jwt.RegisteredClaims
}

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}

func GenerateTokenUsingHs256(key string, expTime time.Duration, user *ent.User) (string, error) {
	claim := MyCustomClaims{
		UserID:   int(user.ID),
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sectran_admin",                                 // 签发者
			Subject:   user.Account,                                    // 签发对象
			Audience:  jwt.ClaimStrings{"SECTRAN"},                     //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expTime)),     //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
			ID:        randStr(10),                                     // wt ID, 类似于盐值
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(key))
	return token, err
}

func ParseTokenHs256(token_string, key string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil //返回签名密钥
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
