/**
* @Author: ZhangZe
* @Date: 2022/4/9 18:47
 */

package common

import (
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
