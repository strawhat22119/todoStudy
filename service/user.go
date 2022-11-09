package service

import (
	"todoStudy/model"
	"todoStudy/pkg/utils"
	"todoStudy/serializer"
)

type UserService struct {
	UserName string `json:"user_name" form:"user_name" binding:"required,min=3,max=8" example:"FanOne"`
	PassWord string `json:"pass_word" form:"pass_word" binding:"required,min=5,max=12" example:"Fan12345"`
}

// 用户注册函数
func (u *UserService) Register() *serializer.Response {
	var user model.User
	var count int64
	//根据user_name 查询表中是否有数据
	model.DB.Model(&model.User{}).Where("user_name=?", u.UserName).First(&user).Count(&count)

	//判断用户是否存在，返回对应json
	if count == 1 {
		return &serializer.Response{
			Msg:    "该用户已存在，注册失败",
			Status: 400,
		}
	}

	//加密密码
	user.UserName = u.UserName
	if err := user.SetPassword(u.PassWord); err != nil {
		return &serializer.Response{
			Msg:    "发生未知错误，请重新注册",
			Status: 400,
		}
	}

	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return &serializer.Response{
			Msg:    "用户创建失败,请重新注册",
			Status: 401,
		}
	}
	return &serializer.Response{
		Msg:    "创建用户成功",
		Status: 200,
	}

}

// 用户登录函数
func (u *UserService) Login() *serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", u.UserName).First(&user).Count(&count)
	if count == 0 {
		return &serializer.Response{
			Msg:    "该用户不存在,请先创建用户",
			Status: 401,
		}
	}
	if !user.CheckPassword(u.PassWord) {
		return &serializer.Response{
			Msg:    "密码不正确",
			Status: 402,
		}
	}
	token, err := utils.GenerateToken(user.ID, u.UserName, 0)
	if err != nil {
		return &serializer.Response{
			Msg:    "登录失败下发token失败",
			Status: 402,
		}
	}

	return &serializer.Response{
		Msg:    "登录成功",
		Data:   "哈哈",
		Status: 200,
		Token:  token,
	}

}
