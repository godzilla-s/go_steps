//上传文件

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		
		defer file.Close()

		fmt.Printf("FileName: %s, Header: %s\n", file, header);
		
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
}

