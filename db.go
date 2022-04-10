package main

import (
	"GinFrame/common"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db  *sql.DB
	err error
)

func InitDB() *sql.DB {
	sourceName := common.GetMysql("./mysql.ini", "dev")
	db, err = sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatal("数据库链接错误！", err)
		return nil
	}
	return db
}
