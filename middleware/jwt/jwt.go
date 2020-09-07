package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//aaa := jwt.Claims{
//UserId: 1,
//}
//ccc,bbb := jwt.CreateToken(&aaa)
//fmt.Println(ccc,bbb)
//ddd,bb := jwt.ValidateToken(ccc)
//fmt.Println(ddd.UserId,bb)

// Claims custom token
type Claims struct {
	UserId int `json:"user_id"` // 用户
	//Version   int32 `json:"version"`    // 版本
	//LoginType int32 `json:"login_type"` // 登录方式
	jwt.StandardClaims
}

// CreateToken create token
func CreateToken(claims *Claims) (signedToken string, success bool) {
	claims.ExpiresAt = time.Now().Add(time.Minute * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	success = true
	return
}

// ValidateToken validate token
func ValidateToken(signedToken string) (claims *Claims, success bool) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected login method %v", token.Header["alg"])
			}
			return []byte("secret"), nil
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
