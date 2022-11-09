package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("AAAA")

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := Claims{
		id,
		username,
		authority,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todoStudy",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SigningString()
}

// 验证用户token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
