package main

import (
	"fmt"
	"time"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id") 
	if err != nil {
		sessionId := createSessionId()
		cookie = &http.Cookie {
			Name : "session_id",
			Value : sessionId,
			HttpOnly: true,
			MaxAge : 10,  //=0 未设置Max-Age属性 <0: 立即删除Max-Age属性; >0 设置Max-Age属性,单位秒 
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

func createSessionId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
