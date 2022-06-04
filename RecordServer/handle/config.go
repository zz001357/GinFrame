/**
* @Author: ZhangZe
* @Date: 2022/6/3 12:33
 */

package handle

import (
	"github.com/spf13/viper"
	"log"
)

func Config() *viper.Viper {
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config_record")
	//添加读取的配置文件路径
	v.AddConfigPath("./RecordServer")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln("配置路径出错", err)
	}
	return v
}
