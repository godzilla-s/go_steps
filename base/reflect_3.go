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

func Abc(nodes ...interface{}) {
	for _, node := range nodes {
		val := reflect.ValueOf(node)
		ind := reflect.Indirect(val)

		if val.Kind() != reflect.Ptr {
			fmt.Println("Args Wrong")
		}
		fmt.Println(ind)
		fmt.Println(ind.Type())

		typ := ind.Type()
		fmt.Println(typ.Kind(), typ.String())

		num := ind.NumField()
		fmt.Println("num: ", num)
		for i:=0; i<num; i++ {
			//fmt.Println(ind.Field(i).Name)
			//val.Elem().Field(i).CanSet()
				
			//fmt.Println(val.Elem().Field(i).CanSet())
			fv := val.Elem().Field(i)
			if fv.Type().Kind() == reflect.String {
				fv.SetString("192.168.1.198")
			}
			
			if fv.Type().Kind() == reflect.Int {
				fv.SetInt(200)
			}
		}
	}
}
func main() {
	d1 := Data {1, "aaaa", true}
	var d2 Data
	CopyStruct(&d1, &d2)
	
	// refelct 映射使用
}
