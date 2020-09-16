package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/server/models"
	"time"
)

const TOKEN_SECRET_KEY = "LOdf%&s465"

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	models.User
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(user models.User) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 12).Unix()),
			Issuer:    "server",
		},
		user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(TOKEN_SECRET_KEY))
	return
}

/**
 * 解析 token
 */
func ParseToken(tokenStr string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return []byte(TOKEN_SECRET_KEY), nil
	})
	claims = token.Claims
	return claims, err
}

func GetTokenVal(c jwt.Claims, key string) (val interface{}) {
	val = c.(jwt.MapClaims)[key]
	return val
}
