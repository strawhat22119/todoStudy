package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	Path       string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		log.Fatalf("读取配置文件失败,请检查路径是否正确")
	}

	LoadServer(file)
	LoadMysql(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
	Path = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", DbUser, DbPassWord, DbHost, DbPort, DbName)

}
