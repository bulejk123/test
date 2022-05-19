package handle

import (
	"easygo/easygo/util"
	"easygo/model/c2s"
	"easygo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	login := c2s.LoginData{}
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	user, err := service.Login(login)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorLoginFail)
	} else {
		expireTime := time.Now().Add(24*time.Hour)
		tokenStr,tokenErr:=service.GenerateToken(user.Id,expireTime)
		if tokenErr == nil {
			c.SetCookie("Authorization",tokenStr,60,"","127.0.0.1",false,true)
		}

		c.JSON(http.StatusOK, user)

	}

}

func AddUser(c *gin.Context) {
	data := c2s.UserData{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	err=service.AddOrUpdateUser(data)
	if err ==nil {
		c.JSON(http.StatusOK,util.Success)
	}else {
		c.JSON(http.StatusBadRequest,util.ErrorAddUserFail)
	}
}

func UpdateUser(c *gin.Context) {
	data := c2s.UserData{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	err=service.AddOrUpdateUser(data)
	if err ==nil {
		c.JSON(http.StatusOK,util.Success)
	}else {
		c.JSON(http.StatusBadRequest,util.ErrorEditUserFail)
	}
}


func DeleteUser(c *gin.Context) {
	data := c2s.DeleteData{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	err=service.DeleteUser(data.Id)
	if err ==nil {
		c.JSON(http.StatusOK,util.Success)
	}else {
		c.JSON(http.StatusBadRequest,util.ErrorDelUserFail)
	}
}
