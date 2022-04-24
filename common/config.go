/**
* @Author: ZhangZe
* @Date: 2022/4/24 19:28
 */

package common

func Connection(key string) string {
	m := make(map[string]string)
	m["go_frame"] = "root:123456@tcp(127.0.0.1:3306)/go_frame?charset=utf8mb4&parseTime=True&loc=Local"
	return m[key]
}
