package jwt

import (
	"../../config"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SECRET = config.Config.Get("jwt.secret").(string)
	//ExpiresAt, _ = time.ParseDuration(config.Config.Get("jwt.ExpiresAt").(string))
)

// Claims custom token
type Claims struct {
	//UserId int `json:"user_id"` // 用户
	UserName string `json:"user_name"` //用户名称
	//Version   int32 `json:"version"`    // 版本
	//LoginType int32 `json:"login_type"` // 登录方式
	jwt.StandardClaims
}

//create token
func CreateToken(claims *Claims) (signedToken string, success bool) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 10).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return
	}
	success = true
	return
}

//validate token
func ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte(SECRET), nil
		})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		success = true
		return
	}

	return
}
