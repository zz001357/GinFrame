package main

import (
	"GinFrame/router"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	/**
	 * @Name 入口文件
	 * @Param
	 * @Return
	 * @Date 2022/4/9 23:57
	 * @Author ZhangZe
	 **/
	gin.DisableConsoleColor()
	folderPath := filepath.Join("./", "log")
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		_ = os.Mkdir(folderPath, os.ModePerm)
	}
	fName := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + fName + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	run := router.Router
	run("192.168.2.135:5000")
}
