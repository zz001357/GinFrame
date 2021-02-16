package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


/*
function=员工登录, methods=["POST"]
interface=/login
name=account info=登录账号 type=string,must=1
name=passwd info=登录密码 type=string,must=1
*/
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "登录",
	})

}
