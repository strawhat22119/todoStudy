package service

import (
	"time"
	"todoStudy/model"
	"todoStudy/pkg/utils"
	"todoStudy/serializer"
)

type TaskService struct {
	Title     string `json:"title" form:"title"`
	Status    int    `json:"status"`
	Content   string `json:"content" form:"content"`
	Uid       int    `json:"uid" `
	StartTime int64
	EndTime   int64
}

// 创建备忘录
func (t *TaskService) CreateTask(token string) *serializer.Response {
	var task model.Task
	//先拉取一下token中的id
	clims, err := utils.ParseToken(token)
	if err != nil {
		return &serializer.Response{
			Status: 400,
			Msg:    "解析token失败",
		}
	}
	id := clims.Id
	var user model.User
	user.ID = id
	if err := model.DB.Model(&model.User{}).First(&user).Error; err != nil {
		return &serializer.Response{
			Status: 400,
			Msg:    "token数据不合法",
		}
	}
	task.Uid = int(user.ID)
	task.Status = 0
	task.StartTime = time.Now().Unix()
	task.Title = t.Title
	task.Content = t.Content
	if err := model.DB.Create(&task).Error; err != nil {
		return &serializer.Response{
			Status: 401,
			Msg:    "创建备忘录失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
		Data:   model.TaskJson(task),
	}
}

// 删除备忘录
func (t *TaskService) DeleteTask(tid string, uid uint) *serializer.Response {
	var task model.Task
	if err := model.DB.Delete(&task, tid).Where("uid=?", uid).Error; err != nil {
		return &serializer.Response{
			Status: 401,
			Msg:    "删除失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}

// 修改备忘录
func (t *TaskService) UpdateTask(tid string, uid uint) *serializer.Response {
	var task model.Task
	if err := model.DB.First(&task, tid).Where("uid=?", uid).Error; err != nil {
		return &serializer.Response{
			Status: 402,
			Msg:    "该用户没有相关备忘录",
		}
	}
	task.Status = t.Status
	task.Title = t.Title
	task.Content = t.Content
	task.Uid = int(uid)
	if err := model.DB.Save(&task).Where("id=?", tid).Error; err != nil {
		return &serializer.Response{
			Status: 403,
			Msg:    "修改备忘录失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "修改备忘录成功",
		Data:   model.TaskJson(task),
	}
}

// 查看备忘录
func (t *TaskService) ShowTask(tid string, uid uint) *serializer.Response {
	var task model.Task
	if err := model.DB.First(&task, tid).Where("uid=?", uid).Error; err != nil {
		return &serializer.Response{
			Status: 401,
			Msg:    "没有查到数据",
		}
	}
	return &serializer.Response{
		Status: 200,
		Msg:    "查询成功",
		Data:   model.TaskJson(task),
	}
}

// 查看所有备忘录
func (t *TaskService) ShowTaskList(uid uint) *serializer.Response {
	var taskList []model.Task
	if err := model.DB.Find(&taskList).Where("uid=?", uid).Error; err != nil {
		return &serializer.Response{
			Status: 401,
			Msg:    "数据查看失败",
		}
	}

	return &serializer.Response{
		Status: 200,
		Data:   model.TaskJsonList(taskList),
		Msg:    "查询成功",
	}
}

// 模糊查询
func (t *TaskService) SearchTask(info string, uid uint) *serializer.Response {
	var task []model.Task
	if err := model.DB.Model(&task).Where("uid=?", uid).Where("title like %?% and content like %?%", info, info).Error; err != nil {
		return &serializer.Response{
			Status: 401,
			Msg:    "查询失败",
		}
	}
	return &serializer.Response{
		Status: 200,
		Data:   model.TaskJsonList(task),
		Msg:    "查询完成",
	}

}
