/**
* @Author: ZhangZe
* @Date: 2022/4/11 10:17
 */

package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"time"
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

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, serviceName string, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: 0, Message: "成功", Data: data})
	apiPath := c.Request.RequestURI
	reqTime := time.Now().Format("2006-01-02 15:04:05")
	respCode := c.Writer.Status()
	reqIp := c.ClientIP()

	db := openSql(Connection().GveMonitor)
	_, err := db.Exec(fmt.Sprintf("insert into t_log(v_service_name,v_api_path,v_req_time,v_resp_code,v_req_ip) values('%s','%s','%s','%d','%s')",
		serviceName, apiPath, reqTime, respCode, reqIp))
	if err != nil {
		log.Println("日志记录出错")
	}
	_ = db.Close()
}

func Failure(c *gin.Context, serviceName string, e error) {
	c.JSON(http.StatusOK, Response{Code: 1, Message: "失败", Data: e})
	apiPath := c.Request.RequestURI
	reqTime := time.Now().Format("2006-01-02 15:04:05")
	respCode := c.Writer.Status()
	reqIp := c.ClientIP()
	errorString := fmt.Sprint(e)
	db := openSql(Connection().GveMonitor)
	sql := "insert into t_log(v_service_name,v_api_path,v_req_time,v_resp_code,v_req_ip,err_content) values(?,?,?,?,?,?)"
	_, errExec := db.Exec(sql, serviceName, apiPath, reqTime, respCode, reqIp, errorString)
	if errExec != nil {
		log.Println("日志记录出错2", errExec)
	}
	_ = db.Close()
}
