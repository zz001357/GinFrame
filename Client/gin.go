/**
* @Author: ZhangZe
* @Date: 2022/5/8 17:05
 */

package main

import (
	"GinFrame/Client/api"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const clientAddr = "192.168.2.135:5000"
const BlogServerAddr = "192.168.2.135:8005"
const ResumeServerAddr = "192.168.2.135:8006"
const PortfoliosServerAddr = "192.168.2.135:8007"

func main() {
	/**
	 * @Name 入口文件
	 * @Param
	 * @Return
	 * @Date 2022/4/9 23:57
	 * @Author ZhangZe
	 **/

	conn1 := conn(BlogServerAddr)
	defer conn1.Close()
	conn2 := conn(ResumeServerAddr)
	defer conn2.Close()
	conn3 := conn(PortfoliosServerAddr)
	defer conn2.Close()
	//设置超时时间
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r := gin.Default()
	r.Use(cors()) //开启中间件 允许使用跨域请求
	folderPath := "./Client/log"
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		_ = os.MkdirAll(folderPath, os.ModePerm)
	}
	fName := time.Now().Format("2006-01-02")
	f, _ := os.Create("./log/" + fName + ".log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	api.BlogView(r, conn1)
	api.ResumeView(r, conn2)
	api.PortfoliosView(r, conn3)

	if err := r.Run(clientAddr); err != nil {
		log.Fatal("程序启动失败:", err)
	}
}

func conn(addr string) *grpc.ClientConn {
	//创建一个grpc连接
	//grpc.WithInsecure()取消明文检测
	conn, errGrpc := grpc.Dial(addr, grpc.WithInsecure())
	if errGrpc != nil {
		panic(errGrpc)
	}
	return conn

}

// 处理跨域

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
