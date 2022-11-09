package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"todoStudy/conf"
)

var DB *gorm.DB

func Database(path string) {
	//判断是否打印日志
	var newLogger logger.Interface
	if conf.AppMode == "debug" {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			logger.Config{
				SlowThreshold:             time.Second,   // 慢 SQL 阈值
				LogLevel:                  logger.Silent, // 日志级别
				IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  false,         // 禁用彩色打印
			},
		)
	}
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明新建表不加s
		},
	})
	//设置数据库链接配置
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Println("db Db err: ", err)
	}
	//设置空闲连接池最大连接数
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)

	//设置连接可复用最大时间
	sqlDb.SetConnMaxLifetime(time.Hour)
	DB = db
	//自动迁移 创建表
	migration()
}
