package routers

import (
	"github.com/gin-gonic/gin"
	"todoStudy/api"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
	}
	return router
}
