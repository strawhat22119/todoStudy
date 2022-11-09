package api

import (
	"github.com/gin-gonic/gin"
	"todoStudy/service"
)

func UserRegister(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		res := userService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(401, "注册失败")
	}
}
func UserLogin(c *gin.Context) {
	var userService service.UserService
	if err := c.ShouldBind(&userService); err != nil {
		res := userService.Login()
		c.JSON(200, res)
	} else {
		c.JSON(401, "登录失败")
	}
}
