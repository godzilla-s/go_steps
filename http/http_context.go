package main

import (
	"fmt"
	"context"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println(ctx)
	fmt.Fprintln(w, ctx)
}

func get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//利用context 存储数据
	ctx = context.WithValue(ctx, "userId", 777)
	ctx = context.WithValue(ctx, "username", "Huason")

	userId := ctx.Value("userId").(int)
	userName := ctx.Value("username").(string)

	fmt.Fprintln(w, userId, userName)
}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/get", get)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println(err)
	}
}
