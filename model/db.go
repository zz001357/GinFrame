package model

import (
	"GinFrame/config"
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB // 所以这个是空
func DBInit() *gorm.DB {
	var mysqlConfig string
	var err error
	dataBase := flag.String("mysql", "dev", "开发环境:dev 生产环境:pro 测试环境:test 默认:dev")
	//执行命令行解析
	flag.Parse()
	//开发环境
	if *dataBase == "dev" {
		mysqlConfig = config.GetMysql("./config/mysql.ini", "dev")
	}
	//测试环境
	if *dataBase == "test" {
		mysqlConfig = config.GetMysql("./config/mysql.ini", "test")
	}
	//生产环境
	if *dataBase == "produce" {
		mysqlConfig = config.GetMysql("./config/mysql.ini", "produce")
	}
	dsn := mysqlConfig
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接错误")
		return nil
	}
	return db
}
