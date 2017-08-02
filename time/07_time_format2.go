package main

import (
  "fmt"
  "time"
)

func time_format(level int) {
  var format string
  switch level {
    case 1:
      format = "2006"
    case 2:
      format = "200601"
    case 3:
      format = "20060102"
    case 4:
      format = "2006010215"
    case 5:
      format = "200601021504"
    default:
      format = "20060102150405"
  }
  fmt.Printf("time: %v\n", time.Now().Format(format))
}

func main() {
  time_format(1)
  time_format(2)
  time_format(3)
  time_format(4)
  time_format(5)
  time_format(6)
}
