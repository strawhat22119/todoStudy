package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"`
	Uid       int    `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:0"` //0:未完成，1:已完成
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

// 序列化
func TaskJson(task Task) Task {
	var t Task
	t.Title = task.Title
	t.Content = task.Content
	t.Status = task.Status
	return t

}

// 序列化
func TaskJsonList(task []Task) []Task {
	var taskList []Task
	for _, v := range task {
		taskList = append(taskList, TaskJson(v))
	}
	return taskList
}
