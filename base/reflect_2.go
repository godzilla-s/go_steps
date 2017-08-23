package main

import (
  "fmt"
  "reflect"
)

type Data struct {
  a   int
  b   string
  c   bool
)

  func attr(d1 interface{}, d2 interface{}){
    // Value 类型
    v_d1 := reflect.ValueOf(d1)
    v_d2 := reflect.ValueOf(d2)
    
    fmt.Println(v_d1, v_d2)
    
    // Type 类型
    t_d1 := reflect.TypeOf(d1)
    t_d2 := reflect.TypeOf(d2)
    
    fmt.Println(t_d1.String(), t_d2.String()) //该类型的值需要多少字节, ~~unsafe.Sizeof
    
    fmt.Println(t_d1, t_d2)
    
    fmt.Println(v_d1.Kind(), v_d2.Kind())
    
    fmt.Println(t_d1.Kind() == reflect.Ptr, t_d2.Kind() == reflect.Ptr)
}
func main() {
  d1 := Data {
    a: 20,
    b: "abcdefg",
    c: true,
  }
  
  d2 := Data {
    a: 100,
    b: "ABCDEFG",
    c: false
  }
  
  attr(d1, &d2)
}
