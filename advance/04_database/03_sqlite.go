package main

import (
	"fmt"
	"database/sql"
	"time"
	"reflect"
	_ "github.com/mattn/go-sqlite3"
)

var (
	create_sql string
	insert_sql string
	query_sql string 
)

type userinfo struct {
	Uid  int
	Username string
	Departname string
	Created  time.Time
}

func createTable(db *sql.DB) {
	create_sql = `CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
		created DATE NULL);`
	db.Exec(create_sql)
}

func insertData(user *userinfo, db *sql.DB) {
	if user == nil {
		fmt.Println("nil userinfo")
		return 
	}

	insert_sql = "INSERT INTO userinfo(username, departname, created) values(?,?,?)"
	stmt, err := db.Prepare(insert_sql)
	if err != nil {
		fmt.Println(err)
		return 
	}

	res, err := stmt.Exec(user.Username, user.Departname, user.Created)
	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println(res)
}

func queryData(data interface{}, db *sql.DB) {
	// query_sql = "SELECT uid,username,departname FROM userinfo"
	query_sql = "SELECT uid,username,departname,created FROM userinfo where uid=?"
	rows, err := db.Query(query_sql, 2)
	if err != nil {
		fmt.Println(err)
		return 
	}
	defer rows.Close()

	//var user userinfo

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
		return
	}

	//var fieldType []reflect.Type 

	
	/*
	var id int
	var name string 
	var depart string 
	
	refs[0] = &id 
	refs[1] = &name 
	refs[2] = &depart 
	*/

	dtype := reflect.ValueOf(data) 
	refs := make([]interface{}, 0, len(cols))
	
	if dtype.Kind() != reflect.Ptr {
		fmt.Println("Not Struct Pointer")
		return 
	}

	indType := reflect.Indirect(dtype)
	if indType.Type().Kind() != reflect.Struct {
		fmt.Println("Not Struct Pointer")
		return 
	}

	columnsMaps := make(map[string]interface{}, len(cols))
	for _, col := range cols {
		var ref interface{}
		columnsMaps[col] = &ref 
		refs = append(refs, &ref)
	}

	for rows.Next() {
		//err = rows.Scan(&user.uid, &user.username, &user.departname, &user.created)
		err = rows.Scan(refs...)
		if err != nil {
			fmt.Println(err)
			return
		}

		break 
		//fmt.Println(id, name, depart)
		//fmt.Println(reflect.TypeOf(refs))
	}

	if len(refs) != indType.NumField() {
		fmt.Println("Fields Not Map")
		return 
	}

	num := indType.NumField()
	for i:=0; i<num; i++ {
		fv := dtype.Elem().Field(i)
		val := reflect.ValueOf(refs[i]).Elem().Interface()
		if fv.Type().Kind() == reflect.String {
			sa := val.([]uint8)
			//s := B2S(sa)
			fv.SetString(fmt.Sprintf("%s", sa))
		}

		if fv.Type().Kind() == reflect.Int {
			a := val.(int64)
			fv.SetInt(a)
			//fmt.Printf("a: %d\n", a)
		}
		//fmt.Println(fv, val)
		//if fv.Type().Kind()
	}
/*
	i := 0
	for k, v := range columnsMaps {
		value := reflect.ValueOf(v).Elem().Interface()
		fmt.Printf("%s: %v", k, value)
		//toStr(value) 
		//setFieldVal(indType.Type().Field(i), value)
		i++
	}
*/
}

/*
func setFieldVal(field reflect.Type, value interface{}) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		if value == nil {
			field.SetInt(0)
		} else {
			field.SetInt(toInt(value))
		}
	case reflect.String:
		if value == nil {
			field.SetString("")
		} else {
			field.SetString(toStr(value))
		}
	}
}
*/
func toStr(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return ""
	}
}

func B2S(bs []int8) string {
    b := make([]byte, len(bs))
    for i, v := range bs {
        b[i] = byte(v)
    }
    return string(b)
}

/*
func toInt(val interface{}) int64 {
	switch v := val.(type) {
	case int,int8,int32,int64:
		return (int64)v 
	default: 
		return 0
	}
}
*/

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	//createTable(db)
	//u := userinfo{username: "hongguofei", departname: "摄影师", created: time.Now(),}
	/*
	insertData(&userinfo{username: "DaFeiNiao", departname: "摄影师", created: time.Now(),}, db)
	insertData(&userinfo{username: "XiaoFeiJi", departname: "工程师", created: time.Now(),}, db)
	insertData(&userinfo{username: "ShiDaNiang", departname: "厨师", created: time.Now(),}, db)
	insertData(&userinfo{username: "EngDaLong", departname: "武术", created: time.Now(),}, db)
	*/
	var user userinfo
	queryData(&user, db)
	fmt.Println(user)
}
