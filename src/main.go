package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	myServer := gin.Default()
	myServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello go server"})
	})
	err := myServer.Run(":8080")
	if err != nil {
		fmt.Println("服务器启动失败！")
	}
}
