package main

import (
	"easygo/config"
	"easygo/model"
	"easygo/router"
)

func main() {
	config.GetGormDb()
	model.InitModel(config.Db)
	router.GinRun()
}
