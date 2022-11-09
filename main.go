package main

import (
	"todoStudy/conf"
	"todoStudy/model"
	"todoStudy/routers"
)

func main() {
	model.Database(conf.Path)
	router := routers.NewRouter()
	router.Run(conf.HttpPort)
}
