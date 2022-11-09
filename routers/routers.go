package routers

import (
	"github.com/gin-gonic/gin"
	"todoStudy/api"
	"todoStudy/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.POST("/user/register", api.UserRegister)
		v1.POST("/user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.Jwt())
		{
			authed.POST("task/create", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("task", api.ShowTaskList)
			authed.PUT("task/Update/:id", api.UpdateTask)
			authed.DELETE("task/Delete/:id", api.DeleteTask)
			authed.POST("task/search/:info", api.SearchTask)
		}
	}
	return router
}
