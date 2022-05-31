/**
* @Author: ZhangZe
* @Date: 2022/4/11 10:17
 */

package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/ini.v1"
	"log"
)

// 根据文件名，段名，键名获取ini的值

func GetIni(filename, expectSection string, key string) string {
	file, err := ini.Load(filename)
	if err != nil {
		log.Println("文件读取错误", err)
		panic(err)

	}
	Port := file.Section(expectSection).Key(key).String()
	return Port
}

func Params(c *gin.Context, key string) string {
	json := make(map[string]interface{})
	_ = c.ShouldBindBodyWith(&json, binding.JSON)
	param := json[key].(string)
	return param
}
