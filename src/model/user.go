package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id            int64     `gorm:"primary_key;column:id"`
	Name          string    `gorm:"column:name"`
	Password      string    `gorm:"column:password"`
	NickName      string    `gorm:"column:nick_name"`
	Addr          string    `gorm:"column:addr"`
	Img           string    `gorm:"column:img"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
	RegisterTime  time.Time `gorm:"column:register_time"`
}

func (User) TableName() string {
	return "user"
}

func init() {
	Migrates = append(Migrates, func(db *gorm.DB) error {
		var err error
		model := &User{}
		if db.HasTable(model) {
			err = db.AutoMigrate(model).Error
		} else {
			err = db.CreateTable(model).Error
		}
		return err
	})
}

type UserRole struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Role     string `gorm:"column:password"`
}

func (UserRole) TableName() string {
	return "user_role"
}

func init() {
	Migrates = append(Migrates, func(db *gorm.DB) error {
		var err error
		model := &UserRole{}
		if db.HasTable(model) {
			err = db.AutoMigrate(model).Error
		} else {
			err = db.CreateTable(model).Error
		}
		return err
	})
}
