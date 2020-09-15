package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TOKEN_SECRET_KEY = "LOdf%&s465"

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid      string
	Username string
}

/**
 * 生成 token
 * SecretKey 是一个 const 常量
 */
func CreateToken(uid, username string) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 12).Unix()),
			Issuer:    "server",
		},
		uid, username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(TOKEN_SECRET_KEY))
	return
}

/**
 * 解析 token
 */
func ParseToken(tokenSrt string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(TOKEN_SECRET_KEY), nil
	})
	claims = token.Claims
	return claims, err
}
