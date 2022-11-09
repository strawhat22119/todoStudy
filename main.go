package main

import (
	"todoStudy/conf"
	"todoStudy/model"
)

func main() {
	model.Database(conf.Path)
}
