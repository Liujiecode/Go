package main

import (
	"fmt"
	"net/http"
	"time"
)

//如何创建服务器？ 如下所示

/*法一
//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!", r.URL.Path)
}
func main() {
	//法一
	http.HandleFunc("/", handler)
	//创建路由
	http.ListenAndServe(":8080", nil)
	//访问 http://localhost:8080/

	//法二、自己创建多路复用器
	mux:=http.MewServeMux()
	mux.HandleFunc("/",handler)
	http.ListenAndServe(":8080",mux)
}
*/
/*法二 自己创建处理器*/
type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过自己创建的处理器处理请求！")
}
func main() {
	/*
		myHandler := MyHandler{}
		http.Handle("/myHandler", &myHandler)
		http.ListenAndServe(":8080", nil)
		//访问  http://localhost:8080/myHandler
	*/

	//升级，详细配置，创建server结构，并配置里面字段
	myHandler := MyHandler{}
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}
	server.ListenAndServe()
}
