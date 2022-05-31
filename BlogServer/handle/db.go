/**
* @Author: ZhangZe
* @Date: 2022/4/11 11:38
 */

package handle

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"runtime"
)

func Connection() string {
	var sqlConfig string
	fmt.Println()
	goos := runtime.GOOS
	if goos == "windows" {
		sqlConfig = Config().GetString("Mysql.dev")
	} else {
		sqlConfig = Config().GetString("Mysql.pro")
	}
	return sqlConfig
}

func openSql(dbConnection string) *sql.DB {
	db, sqlErr := sql.Open("mysql", dbConnection)
	if sqlErr != nil {
		log.Println("数据库连接出错！")
		panic(sqlErr)
	}
	return db
}

//查询

func ReadSql(sql string, dbConnection string) ([]string, []map[string]string, error) {
	var records []map[string]string
	var columns []string //数据库列
	db := openSql(dbConnection)
	data, queryErr := db.Query(sql)
	if queryErr != nil {
		log.Println("查询错误！", queryErr)
		return columns, records, queryErr
	}
	columns, _ = data.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	//先给scanArgs 分配地址，scanArgs的每个值指向values相应值的地址
	//scanArgs 的作用是存放扫描的结果，因为扫描需要用的是地址
	for i := range values {
		scanArgs[i] = &values[i]
	}
	//获得数据
	for data.Next() {
		record := make(map[string]string)
		_ = data.Scan(scanArgs...)
		//因为这里的values的每个值的地址对应的是scanArgs
		for i, col := range values {
			if col != nil {
				//将byte数据转换成字符串，对应列表顺序，存入字典record
				record[columns[i]] = string(col.([]byte))
			}
		}
		records = append(records, record)
	}
	_ = db.Close()
	return columns, records, nil
}

//插入

func InsertSql(sql string, dbConnection string) error {
	db := openSql(dbConnection)
	_, err := db.Exec(sql)
	if err != nil {
		log.Println(err)
		return err
	}
	_ = db.Close()
	return nil
}
