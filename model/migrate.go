package model

// 迁移
func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}, &Task{})
}
