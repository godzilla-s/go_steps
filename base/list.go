// container/list 实现了双向链表

package main

import (
	"fmt"
	"container/list"
)

func list_print(lst *list.List) {
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	println()
}

func list_get(lst *list.List, pos int) *list.Element {
	size := lst.Len()
	if pos >= size {
		return lst.Back()
	}

	if pos <= 0 {
		return lst.Front()
	}

	e := lst.Front()
	for i:=1; i<pos; i++ {
        e = e.Next()
    }

	return e
}

func main() {
	lst := list.New()

	lst.PushFront(10)
	lst.PushFront(13)
	lst.PushFront("Gee")
	
	fmt.Printf("Length: %d\n", lst.Len())
	list_print(lst)

	lst.PushBack("Toe")
	list_print(lst)

	lst.InsertAfter(50, list_get(lst, 2))
	list_print(lst)
}
