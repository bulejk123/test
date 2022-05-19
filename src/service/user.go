package service

import (
	"easygo/config"
	"easygo/model"
	"easygo/model/c2s"
	"log"
	"time"
)

func Login(data c2s.LoginData) (*model.User, error) {
	user := &model.User{}
	err := config.Db.Where("name = ? and password = ?", data.Name, data.Password).First(user).Error
	if err != nil {
		log.Println("find user fail ", err)
		return nil, err
	}
	err=UpdateLoginTime(user.Id)
	if err == nil {
		user.LastLoginTime = time.Now()
	}
	return user, err
}

func UpdateLoginTime(id int64)error{
	err := config.Db.Model(model.User{}).Update("last_login_time=?", time.Now()).Where("id = ?", id).Error
	if err != nil {
		log.Println("update user loginTime fail")
		return err
	}
	return nil
}

func AddOrUpdateUser(data c2s.UserData) error {
	user := model.User{
		Name:         data.Name,
		Password:     data.Password,
		NickName:     data.NickName,
		Img:          data.Img,
		Addr:         data.Addr,
		RegisterTime: time.Now(),
	}
	err := config.Db.Save(&user).Error
	return err
}

//func UpdateUser(data c2s.UserData) error {
//	user := model.User{
//		Id:           data.Id,
//		Name:         data.Name,
//		Password:     data.Password,
//		NickName:     data.NickName,
//		Img:          data.Img,
//		Addr:         data.Addr,
//		RegisterTime: time.Now(),
//	}
//	db := config.Db
//	err := db.Save(&user).Error
//	return err
//}

func DeleteUser(id int64) error {
	err := config.Db.Where("id = ? ", id).Delete(model.User{}).Error
	return err
}
