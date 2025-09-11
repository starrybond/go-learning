package utils

// 这段代码是整套系统的JWT工具箱，只做三件小事：
// 生成登录令牌（GenToken）
// 解析/校验令牌（ParseToken）
// 携带数据（用户 ID + 用户名 + 过期时间）

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("BLOG_SECRET")

type Claims struct {
	UserID uint   `json:"id"`
	Name   string `json:"username"`
	jwt.RegisteredClaims
}

func GenToken(uid uint, username string) (string, error) {
	c := Claims{
		UserID: uid,
		Name:   username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(secret)
}

func ParseToken(tokenString string) (*Claims, error) {
	var c Claims
	token, err := jwt.ParseWithClaims(tokenString, &c, func(t *jwt.Token) (interface{}, error) {
		// 确认签名算法
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid { // ← 关键：必须检查 Valid
		return nil, jwt.ErrSignatureInvalid
	}
	return &c, nil
}
