/**
* @Author: ZhangZe
* @Date: 2022/4/24 19:28
 */

package common

type connection struct {
	GoFrame string
}

func Connection() *connection {
	con := new(connection)
	con.GoFrame = "root:123456@tcp(127.0.0.1:3306)/go_frame?charset=utf8mb4&parseTime=True&loc=Local"
	return con
}
