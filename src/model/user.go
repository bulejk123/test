package model

type User struct {
	Id int64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}
