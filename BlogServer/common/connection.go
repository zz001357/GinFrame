/**
* @Author: ZhangZe
* @Date: 2022/5/9 15:47
 */

package common

import "runtime"

func Connection() string {
	var sqlConfig string
	sysType := runtime.GOOS
	if sysType == "windows" {
		sqlConfig = Config().GetString("Mysql.dev")
	} else {
		sqlConfig = Config().GetString("Mysql.pro")
	}
	return sqlConfig
}
