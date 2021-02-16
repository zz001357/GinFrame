package main

import (
	"GinFrame/model"
	"GinFrame/router"
	"github.com/gin-gonic/gin"
	"log"
)

//获取环境变量
func main() {
	r := gin.Default()
	router.InitRouter(r)
	model.DBInit()
	if err := r.Run(":9999"); err != nil {
		log.Fatal("程序启动失败:", err)
	}
}
