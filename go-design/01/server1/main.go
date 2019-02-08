// server1 是一个简单响应服务器

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/",handler) // 响应请求调用处理函数
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}