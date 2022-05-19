package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	user     = "root"
	password = "Test1234"
	host     = "127.0.0.1"
	port     = 3306
	db       = "test"
)

var (
	Db *gorm.DB
)

func GetGormDb() {
	var err error
	Db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local&charset=utf8mb4,utf8",
		user, password, host, port, db))
	if err != nil {
		log.Println("db fail",err)
	}
}
