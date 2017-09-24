package main

import (
  "fmt"
  "reflect"
)

type Data struct {
	a     int		`test: a`
	b     string		`test: b`
	c     bool		`test: c`
}

func CopyStruct(src, dst interface{}) {
	vsrc := reflect.ValueOf(src)
	vdst := reflect.ValueOf(dst)

	if vsrc.Kind() != reflect.Ptr || vdst.Kind() != reflect.Ptr {
		fmt.Println("Not Struct Pointer")
		return
	}

	numSrcFld := vsrc.NumField()
	numDstFld := vdst.NumField()

	if numSrcFld != numDstFld {
		fmt.Println("Number of field not map")
		return
	}

	t1 := reflect.TypeOf(src)
	if t.Kind() == reflect.Struct {
		fmt.Println(t1.Field(i).Name, t1.Filed(i).Type, t1.Field(i).Type)
	}
	
	for i:=0; i<numSrcFld; i++ {
		if vdst.Elem().Field(i).CanSet() {
			vdst.Elem().Field(i).Set(vsrc.Elem().Field(i))
		}
	}
}

func main() {
	d1 := Data {1, "aaaa", true}
	var d2 Data
	CopyStruct(&d1, &d2)
	
	// refelct 映射使用
}
