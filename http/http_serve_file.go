package main

import(
	"io"
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/cat1", cat1)
	http.HandleFunc("/cat2", cat2)
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	//下面方式图片不会显示到html上
	io.WriteString(w, `<img src="toby.jpg"> `)
}

//下面方式可以传递图片
func cat1(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("img/cat01.jpg")
	if err != nil {
		io.WriteString(w, "Not Found Image")
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

// 下面使用http.Error() 返回错误
// http.ServeContent() 返回页面
func cat2(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("img/cat04.jpg")
	if err != nil {
		//io.WriteString(w, "Not Found Image")
		http.Error(w, "File Not Found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		//io.WriteString(w, "Not Found Image")
		http.Error(w, "File Not Found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}

//http.ServeFile() 使用
func dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "img/dog.jpg")
}
