// 事务处理

package main

import (
  "log"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

var (
  dbname = "root"
  dbpwd = "root123456"
  dbuser = "mybase"
  dbip = "127.0.0.1"
  dbport = "3306"
)

// 提交事务
func db_transaction_commit(db *sql.DB) {
  tx, err := db.Begin()
	if err != nil {
		log.Fatal("Prepare err:", err)
	}

	sql := "INSERT INTO mytable(id, name, email) values(?,?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal("Prepare err:", err)
	}

	_, err = stmt.Exec(6, "Gergeo", "franc@tmerik.com")
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
    log.Fatal("Commit: ", err)
  }
}

// 回滚事务
func db_transaction_rollback(db *sql.DB) {
  tx, err := db.Begin()
	if err != nil {
		log.Fatal("Prepare err:", err)
	}

	sql := "INSERT INTO mytable(id, name, email) values(?,?,?)"
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal("Prepare err:", err)
	}

	_, err = stmt.Exec(6, "Gergeo", "franc@tmerik.com")
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Rollback(); err != nil {
    log.Fatal("Rollback: ", err)
  }
}

func db_transaction_insert(db *sql.DB) {
  connstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbname, dbpwd, dbip, dbport, dbuser)
  db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal("Open MySQL fail:", err)
	}
	defer db.Close()
	log.Println("Open MySQL OK")
  
  db_transaction_commit(db)
  db_transaction_rollback(db)
}
