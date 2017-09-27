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

func queryOne(data interface{}, db *sql.DB) {
	vType := reflect.ValueOf(data)
	if vType.Kind() != reflect.Ptr {
		fmt.Println("Not Point Struct")
		return 
	}
	indType := reflect.Indirect(vType)
	if indType.Kind() != reflect.Struct {
		fmt.Println("Not Point Struct")
		return 
	}

	numFld := indType.NumField()
	fmt.Println(numFld)
	refs := make([]interface{}, 0, numFld)
	for i:=0; i<numFld; i++ {
		var ref interface{}
		refs = append(refs, &ref)
	}

	query_sql = "SELECT uid,username,departname,created FROM userinfo"
	err := db.QueryRow(query_sql, 2).Scan(refs...)
	if err != nil {
		fmt.Println("scan:", err)
		return 
	}
	mapping2(vType.Elem(), refs)
}

func queryAll(data interface{}, db *sql.DB) {
	var dataType reflect.Value 

	vtype := reflect.ValueOf(data)
	indtype := reflect.Indirect(vtype)

	if vtype.Kind() != reflect.Ptr && indtype.Kind() != reflect.Slice {
		fmt.Println("Not struct field")
		return 
	}

	dataType = indtype 

	query_sql = "SELECT uid,username,departname,created FROM userinfo"
	rows, err := db.Query(query_sql)
	if err != nil {
		fmt.Println(err)
		return 
	}

	etype := indtype.Type().Elem()	
	numFld := etype.NumField()
	refs := make([]interface{}, 0, numFld)
	for i:=0; i<numFld; i++ {
		var ref interface{}
		refs = append(refs, &ref)
	}

	dataSet := dataType
	count := 0
	for rows.Next() {
		err := rows.Scan(refs...)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := reflect.New(etype)
		if data.Kind() == reflect.Ptr {
			data = data.Elem()
		}
		mapping2(data, refs)
		dataSet = reflect.Append(dataSet, data)
		count++
	}

	if count > 0 {
		dataType.Set(dataSet)
	}
}

func mapping2(vtype reflect.Value, refs []interface{}) {
	//vtype := reflect.ValueOf(data)
	num := len(refs)
	for i:=0; i<num; i++ {
		fv := vtype.Field(i)
		if !fv.CanSet() {
			continue 
		}

		val := reflect.ValueOf(refs[i]).Elem().Interface()
		if fv.Type().Kind() == reflect.String {
			str := val.([]uint8)
			fv.SetString(fmt.Sprintf("%s", str))
		}

		if fv.Type().Kind() == reflect.Int {
			data := val.(int64)
			fv.SetInt(data)
		}
	}
}

func mapping(data interface{}, refs []interface{}) {
	vtype := reflect.ValueOf(data)
	num := len(refs)
	for i:=0; i<num; i++ {
		fv := vtype.Elem().Field(i)
		if !fv.CanSet() {
			continue 
		}

		val := reflect.ValueOf(refs[i]).Elem().Interface()
		if fv.Type().Kind() == reflect.String {
			str := val.([]uint8)
			fv.SetString(fmt.Sprintf("%s", str))
		}

		if fv.Type().Kind() == reflect.Int {
			data := val.(int64)
			fv.SetInt(data)
		}
	}
}

func queryData(data interface{}, db *sql.DB) {
	// query_sql = "SELECT uid,username,departname FROM userinfo"
	query_sql = "SELECT uid,username,departname,created FROM userinfo"
	rows, err := db.Query(query_sql)
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

	dtype := reflect.ValueOf(data) 
	refs := make([]interface{}, 0, len(cols))

	/*
	var id int
	var name string 
	var depart string 
	
	refs[0] = &id 
	refs[1] = &name 
	refs[2] = &depart 
	*/

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
		fmt.Println(col)
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

	numFld := indType.NumField()
	for i:=0; i<numFld; i++ {
		fv := dtype.Elem().Field(i)
		if !fv.CanSet() {
			continue 
		}
		val := reflect.ValueOf(refs[i]).Elem().Interface()
		if fv.Type().Kind() == reflect.String {
			str := val.([]uint8)
			fv.SetString(fmt.Sprintf("%s", str))
		}

		if fv.Type().Kind() == reflect.Int {
			data := val.(int64)
			fv.SetInt(data)
		}
	}

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

func toInt(val interface{}) int {
	switch v := val.(type) {
	case int,int8,int32,int64:
		return v 
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
	//insertData(&userinfo{username: "DaFeiNiao", departname: "摄影师", created: time.Now(),}, db)
	var user userinfo
	//queryData(&user, db)
	
	
	queryOne(&user, db)
	fmt.Println(user)
	
	var users []userinfo
	queryAll(&users, db)
	fmt.Println(users)
}
