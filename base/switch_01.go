import "fmt"

const (
    Mon = "星期一"
    Tue = "星期二"
    Wen = "星期三"
    Thu = "星期四"
    Fri = "星期五"
    Sat = "星期六"
    Sun = "星期日"
    Unknow = "未知"
)

func switch_case(a int) {
    switch {
        case a == 1:
            fmt.Println(Wen)
        case a == 2:
            fmt.Println(Tue)
        case a == 3:
            fmt.Println(Wen)
        case a == 4:
            fmt.Println(Thu)
        case a == 5:
            fmt.Println(Fri)
        case a == 6:
            fmt.Println(Sat)
        case a == 7:
            fmt.Println(Sun)
        default:
            fmt.Println(Unknow)
     }
 }
 
 func switch_case_without_fallthrough() {
    fmt.Println("------switch case without fallthrough")
    str := "StringString"
    switch {
        case len(str) >= 10:
            fmt.Println("length of str is more 10 bytes")
        case str == "StringStr":
            fmt.Println("equal StringStr")
        case str == "StringString":
            fmt.Println(str)
    }
}

func swicth_case_with_fallthrough() {
    fmt.Println("------swicth with fallthrough")
    str := "StringString"
    switch {
        case len(str) >= 10:
            fmt.Println("length of str is more 10 bytes")
            fallthrough
        case str == "StringStr":
            fmt.Println("equal StringStr")
        case str == "StringString":
            fmt.Println(str)
    }
}

func main() {
    switch_case(5)
    switch_case(3)
    switch_case(10)

    switch_case_without_fallthrough()

    swicth_case_with_fallthrough()
}

    
