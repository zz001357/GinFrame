package main

import (
	"GinFrame/router"
)

func init() {
	InitDB() //数据库连接
}

func main() {
	/**
	 * @Name 入口文件
	 * @Param
	 * @Return
	 * @Date 2022/4/9 23:57
	 * @Author ZhangZe
	 **/
	run := router.Router
	run(":8000")
}
