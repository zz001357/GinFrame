/**
* @Author: ZhangZe
* @Date: 2022/4/24 19:28
 */

package common

type connection struct {
	//GoFrame         string
	GveServiceCenter string
	GveBlog          string
	GveLife          string
	GveResume        string
	GveMonitor       string
}

func Connection() *connection {
	con := new(connection)
	//con.GoFrame = "root:123456@tcp(127.0.0.1:3306)/go_frame?charset=utf8mb4&parseTime=True&loc=Local"
	con.GveServiceCenter = "root:123456@tcp(127.0.0.1:3306)/gve_service_center?charset=utf8mb4&parseTime=True&loc=Local"
	con.GveBlog = "root:123456@tcp(127.0.0.1:3306)/gve_blog?charset=utf8mb4&parseTime=True&loc=Local"
	con.GveLife = "root:123456@tcp(127.0.0.1:3306)/gve_life?charset=utf8mb4&parseTime=True&loc=Local"
	con.GveResume = "root:123456@tcp(127.0.0.1:3306)/gve_resume?charset=utf8mb4&parseTime=True&loc=Local"
	con.GveMonitor = "root:123456@tcp(127.0.0.1:3306)/gve_monitor?charset=utf8mb4&parseTime=True&loc=Local"
	return con
}
