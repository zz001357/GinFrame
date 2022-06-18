/**
* @Author: ZhangZe
* @Date: 2022/5/8 17:05
 */

package main

import (
	"GinFrame/Client/common"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/host"
	"io"
	"log"
	"os"
	"time"
)

const clientAddr = ":5000"

func main() {
	/**
	 * @Name 入口文件
	 * @Param
	 * @Return
	 * @Date 2022/4/9 23:57
	 * @Author ZhangZe
	 **/
	nInfo, _ := host.Info()
	gin.SetMode(gin.ReleaseMode)

	if nInfo.OS != "windows" {
		gin.SetMode(gin.ReleaseMode)

	}
	r := gin.Default()
	r.Use(common.Cors())       //开启中间件 允许使用跨域请求
	r.Use(common.TlsHandler()) //开启中间件 使用https

	folderPath := "./Client/log"
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(folderPath, os.ModePerm)
	}
	fName := time.Now().Format("2006-01-02")
	f, _ := os.Create(folderPath + "/" + fName + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	conn1, conn2, conn3, conn4 := Server(r)
	if nInfo.OS == "windows" {
		if err := r.Run(clientAddr); err != nil {
			log.Fatal("程序启动失败:", err)
		}
	} else {
		if err := r.RunTLS(clientAddr, "ggva.ren_bundle.pem", "ggva.ren.key"); err != nil {
			log.Fatal("程序启动失败:", err)
		}

	}
	_ = conn1.Close()
	_ = conn2.Close()
	_ = conn3.Close()
	_ = conn4.Close()
}
