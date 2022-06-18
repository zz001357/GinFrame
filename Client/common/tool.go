/**
* @Author: ZhangZe
* @Date: 2022/4/11 10:17
 */

package common

import (
	"encoding/json"
	"fmt"
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

func IpUntil() (string, string) {
	var returnOut string
	responseClient, errClient := http.Get("http://ip.dhcp.cn/?ip") // 获取外网 IP
	if errClient != nil {
		log.Println("获取外网 IP 失败，请检查", errClient)
		returnOut = "获取外网 IP 失败"
		return returnOut, ""
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()
	body, _ := ioutil.ReadAll(responseClient.Body)
	returnOut = fmt.Sprintf("%s", string(body))
	resp, err := http.Get("http://whois.pconline.com.cn/ipJson.jsp?json=true&level=3&ip=" + returnOut)
	if err != nil {
		log.Println("解析ip错误", err)
	}
	defer resp.Body.Close()
	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	addrBody, _ := ioutil.ReadAll(decoder.NewReader(resp.Body))
	addrMap := make(map[string]interface{})
	_ = json.Unmarshal(addrBody, &addrMap)
	addr := addrMap["addr"].(string)

	return returnOut, addr
}
