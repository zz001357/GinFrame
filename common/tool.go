/**
* @Author: ZhangZe
* @Date: 2022/4/11 10:17
 */

package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"log"
)

// 根据文件名，段名，键名获取ini的值

func GetMysql(filename, expectSection string) string {
	file, err := ini.Load(filename)
	if err != nil {
		log.Println("文件读取错误", err)
		panic(err)

	}
	Port := file.Section(expectSection).Key("dsn").String()
	return Port
}

func Params(c *gin.Context, key string) string {
	json := make(map[string]interface{})
	_ = c.BindJSON(&json)
	param := json[key].(string)
	return param
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) Result(code int, message string, data interface{}) Response {
	r.Code = code
	r.Message = message
	r.Data = data
	return r
}
