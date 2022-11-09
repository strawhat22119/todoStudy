package api

import (
	"github.com/gin-gonic/gin"
	"todoStudy/pkg/utils"
	"todoStudy/service"
)

// 创建备忘录
func CreateTask(c *gin.Context) {
	var taskService service.TaskService
	if err := c.ShouldBind(&taskService); err != nil {
		c.JSON(401, "注册失败")
	} else {
		token := c.GetHeader("Authorization")
		res := taskService.CreateTask(token)
		c.JSON(200, res)
	}
}

// 查看备忘录
func ShowTask(c *gin.Context) {
	var taskService service.TaskService
	if id := c.Param("id"); id == "" {
		c.JSON(401, "id为空")
	} else {
		token, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := taskService.ShowTask(id, token.Id)
		c.JSON(200, res)
	}
}

// 查看所有备忘录
func ShowTaskList(c *gin.Context) {
	var taskService service.TaskService
	token, _ := utils.ParseToken(c.GetHeader("Authorization"))
	res := taskService.ShowTaskList(token.Id)
	c.JSON(200, res)
}

// 修改备忘录
func UpdateTask(c *gin.Context) {
	var taskService service.TaskService
	if err := c.ShouldBind(&taskService); err != nil {
		c.JSON(401, "修改请求错误")
	} else {
		if id := c.Param("id"); id == "" {
			c.JSON(401, "id不能为空")
		} else {
			token, _ := utils.ParseToken(c.GetHeader("Authorization"))
			res := taskService.UpdateTask(id, token.Id)
			c.JSON(200, res)
		}
	}
}

//删除备忘录

func DeleteTask(c *gin.Context) {
	var taskService service.TaskService
	if id := c.Param("id"); id == "" {
		c.JSON(401, "id不能为空")
	} else {
		token, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := taskService.DeleteTask(id, token.Id)
		c.JSON(200, res)
	}
}

// 模糊查询
func SearchTask(c *gin.Context) {
	var taskService service.TaskService
	if info := c.Param("info"); info == "" {
		c.JSON(401, "查询条件为空")
	} else {
		token, _ := utils.ParseToken(c.GetHeader("Authorization"))
		res := taskService.SearchTask(info, token.Id)
		c.JSON(200, res)
	}
}
