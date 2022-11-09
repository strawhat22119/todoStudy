package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PasswordDigest string //储存的是加密后的密文密码
}
