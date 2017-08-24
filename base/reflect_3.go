package main

import (
  "fmt"
  "reflect"
)

type Data struct {
  a     int
  b     string
  c     bool
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

	for i:=0; i<numSrcFld; i++ {
		if vdst.Elem().Field(i).CanSet() {
			vdst.Elem().Field(i).Set(vsrc.Elem().Field(i))
		}
	}
}

func main() {
	d1 := Data {1, "aaaa", true}
	
	
}
