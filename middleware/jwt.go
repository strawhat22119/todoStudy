package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"todoStudy/pkg/utils"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 403
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 401
			}
		}
		fmt.Println("token:", token)
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "Token解析失败",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}
