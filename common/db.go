/**
* @Author: ZhangZe
* @Date: 2022/4/11 11:38
 */

package common

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB
var err error

func init() {
	sourceName := GetMysql("./mysql.ini", "dev")
	db, err = sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatal("数据库链接错误！", err)
		return
	}
}

//查询

func ReadSql(sql string) (map[string]string, error) {
	record := make(map[string]string)
	data, err := db.Query(sql)
	if err != nil {
		return record, err
	}
	columns, _ := data.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	//先给scanArgs 分配地址，scanArgs的每个值指向values相应值的地址
	//scanArgs 的作用是存放扫描的结果，因为扫描需要用的是地址
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for data.Next() {
		err = data.Scan(scanArgs...)
		//因为这里的values的每个值的地址对应的是scanArgs
		for i, col := range values {
			if col != nil {
				//将byte数据转换成字符串，对应列表顺序，存入字典record
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	return record, err
}
