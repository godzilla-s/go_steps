package main

import (
	"fmt"
	"database/sql"
	"reflect"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

type DBHandle struct {
	db 			*sql.DB 
	driver   	string
	dblink 		string
	table		tableInfo
}

type fieldInfo struct {
	fldName 	string
	fldType 	reflect.Type
	fldTag		reflect.StructTag 
}

type tableInfo struct {
	tblName 	string
	fields 		[]fieldInfo
}

func getTableInfo(data interface{}) (tableInfo, error) {
	dtype := reflect.TypeOf(data)

	var table tableInfo
	if dtype.Kind() != reflect.Struct {
		return tableInfo{}, fmt.Errorf("Not Struct type")
	}

	numFld := dtype.NumField()
	for i:=0; i<numFld; i++ {
		field := fieldInfo{dtype.Field(i).Name, dtype.Field(i).Type, dtype.Field(i).Tag}
		table.fields = append(table.fields, field)
	}

	table.tblName = dtype.Name()
	return table, nil
}

const (
	DB_INSERT = 1
	DB_SELECT = 2
	DB_UPDATE = 3
)

func parseSQL(oper int, table tableInfo) string {
	var (
		outsql 	string 
		fldNames []string 
		placeHolder []string
	)

	for _, v := range table.fields {
		fldNames = append(fldNames, v.fldName)
		placeHolder = append(placeHolder, "?")
	}

	if oper == DB_INSERT {
		outsql = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", table.tblName, strings.Join(fldNames, ","), strings.Join(placeHolder, ","))
	} 

	if oper == DB_UPDATE {
		outsql = "update"
	}

	return outsql
}

func NewDBHandle(driver, link string) *DBHandle {
	return &DBHandle{
		driver: driver,
		dblink: link,
	}
}

func (handle *DBHandle) Open() (*DBHandle, error) {
	db, err := sql.Open(handle.driver, handle.dblink)
	if err != nil {
		return handle, err
	}
	handle.db = db
	return handle, nil 
}

func (handle *DBHandle) New(data interface{}) error {
	table, err := getTableInfo(data)
	if err != nil {
		return err
	}

	handle.table = table

	return nil
}

type DB_UserInfo struct {
	id 			int 
	name 		string 
	title 		string 
	birth 		string 
}

func (handle *DBHandle) Save() {
	sql := parseSQL(DB_INSERT, handle.table)
	fmt.Println(sql)
}

func main() {
	handle, err := NewDBHandle("sqlite3", "test.db").Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	var user DB_UserInfo

	err = handle.New(user)
	if err != nil {
		fmt.Println(err)
		return 
	}

	handle.Save()
}
