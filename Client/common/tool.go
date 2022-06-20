/**
* @Author: ZhangZe
* @Date: 2022/4/11 10:17
 */

package common

import (
	"encoding/json"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
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

func IpUntil(c *gin.Context) (string, string) {
	ip := c.ClientIP()
	resp, err := http.Get("http://whois.pconline.com.cn/ipJson.jsp?json=true&level=3&ip=" + ip)
	if err != nil {
		log.Println("解析ip错误", err)
	}
	defer resp.Body.Close()
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	addrBody, _ := ioutil.ReadAll(decoder.NewReader(resp.Body))
	addrMap := make(map[string]interface{})
	_ = json.Unmarshal(addrBody, &addrMap)
	addr := addrMap["addr"].(string)
	return ip, addr
}
