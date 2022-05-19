package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	Migrates []func(db *gorm.DB) error
)

func InitModel(db *gorm.DB)  {
	//fmt.Print(Migrates)
	//fmt.Printf("%v",Migrates)
	for _, f := range Migrates {
		if err := f(db); err != nil {
			log.Fatal("init db failed: ", err.Error())
			return
		}
	}
}