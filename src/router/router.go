package router

import (
	"easygo/handle"
	"easygo/service"
	"github.com/gin-gonic/gin"
	"log"
)

func GinRun(){
	router:=gin.Default()
	router.POST("/api/v1/user/login", handle.Login)
	router.Use(service.Cors())
	router.Use(service.JWTAuth())
	r:=router.Group("/api/v1/user")
	{
		r.POST("/add", handle.AddUser)
		r.POST("/edit", handle.UpdateUser)
		r.POST("/del", handle.DeleteUser)
	}
	//r1:=router.Group("/api/v1/tr")
	{
		//r1.POST("/tr_data", handle.GetTrData)

	}
	err:=router.Run(":9001")
	if err!=nil {
		log.Fatal("服务器启动失败")
	}

}
