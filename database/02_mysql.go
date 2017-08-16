// 连接池

package main

import (
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var (
  dbname = "root"
  dbpwd = "root123456"
  dbuser = "mybase"
  dbip = "127.0.0.1"
  dbport = "3306"
)

func init() {
	conninfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbname, dbpwd, dbip, dbport, dbuser)
  dbconn, err := sql.Open("mysql", conninfo)
	if err != nil {
		log.Fatal("Open Mysql:", err)	
	}
  dbconn.SetMaxOpenConns(1000)
  // 如果保持一个空闲链接太长，则会出现出现网络连接的问题，类似下面错误： 
	// unexpected EOF
	// write tcp 192.168.3.90:3306: broken pipe
  // 可以试着：db.SetMaxIdleConns(0)
  dbconn.SetMaxIdleConns(500)
  dbconn.Ping()
  
  db = dbconn
}

func db_query1() {

}

func db_query2() {

}

func db_insert() {

}

func db_update() {

}

func main() {
  return 
}
