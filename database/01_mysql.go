// for test
// 需要安装mysql数据库，并创建mybase库， mytable表
// 简单的mysql数据库使用

package main

import (
  "fmt"
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// insert database
func db_insert(db *sql.DB) {
	sql_01 := "INSERT INTO mytable(id, name, email) values(?, ?, ?)"
	stmt, err := db.Prepare(sql_01)
	if err != nil {
		log.Fatal("db Prepare fatal\n")
	}

	defer stmt.Close()

	_, err = stmt.Exec(2, "Bolom", "bolom@tmerik.com")
	if err != nil {
		log.Println("Exec error: ", err)
	}

	_, err = stmt.Exec(3, "Caler", "caler@tmerik.com")
	if err != nil {
		log.Println("Exec error: ", err)
	}
}

// 更新数据库
func db_update(db *sql.DB) {
	sql_03 := "UPDATE mytable SET email=? where id=?"
	stmt, err := db.Prepare(sql_03)
	if err != nil {
		log.Fatal("Update Prepare error:", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec("XXX@undefine.com", 3)
	if err != nil {
		log.Println("Updare Exec err:", err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root123456@/mybase?charset=utf8&loc=Local")
	if err != nil {
		log.Fatal("Open MySQL fail:", err)
	}
	defer db.Close()
	log.Println("Open MySQL OK")

	if err := db.Ping(); err != nil {
		log.Fatal("Ping err:", err)
	}

	db_insert(db)

	db_update(db)

	//query
	sql_02 := "SELECT id, name, email FROM mytable"
	result, err := db.Query(sql_02)
	if err != nil {
		log.Fatal("Query fail: ", err)
	}
	defer result.Close()

	var id  int
	var name string
	var email  string

	for result.Next() {
		if err := result.Scan(&id, &name, &email); err != nil {
			log.Fatal("Get Scan err: ", err)
		}

		log.Println("Id:", id, " Name:", name, "Email:", email)
	}
}
