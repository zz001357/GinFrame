/**
* @Author: ZhangZe
* @Date: 2022/4/24 19:28
 */

package handle

import (
	"github.com/spf13/viper"
	"log"
)

func Config() *viper.Viper {
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config_resume")
	//添加读取的配置文件路径
	v.AddConfigPath("./ResumeServer")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln("配置路径出错", err)
	}
	return v
}
