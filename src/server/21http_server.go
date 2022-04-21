package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// /*server(服务端)*/

// /*简单html页面搭建*/
// func he(w http.ResponseWriter, r *http.Request) {
// 	// str := "<h1>你好 Go</h1>"
// 	// //或者打开文本,注意文本位置与main.py相对位置
// 	// str2, err := ioutil.ReadFile("../go_learn/xx.txt")
// 	// if err != nil {
// 	// 	w.Write([]byte(fmt.Sprintf("error%v", err)))
// 	// }
// 	// w.Write(str2)
// 	//直接打开html
// 	str3, err := ioutil.ReadFile("../go_learn/xx.html")
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf("error%v", err)))
// 	}
// 	w.Write(str3)
// }
// func client(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf(r.Host)
// 	fmt.Println(r.Body)
// 	fmt.Println(r.URL)
// 	fmt.Println(r.Response)
// 	fmt.Println(ioutil.ReadAll(r.Body)) //我在服务端打印客户端发来的body []，get方法的body为空
// 	fmt.Println(r.URL.Query())          //获得url传递的参数  map[age:[18] name:[sb]]
// 	name := r.URL.Query().Get("name")
// 	age := r.URL.Query().Get("age")
// 	fmt.Printf("Get()传递的参数name： %v ,age: %v", name, age)
// 	w.Write([]byte("ok"))
// }

// func main() {

// 	http.HandleFunc("/index/", he)
// 	http.HandleFunc("/client/", client)

// 	http.ListenAndServe("127.0.0.1:9090", nil) //监听

// }

// /*设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。*/
// func Http_02() {
// 	/*Get请求*/
// 	resp, err := http.Get("https://www.baidu.com")
// 	if err != nil {
// 		fmt.Printf("Get failed,err:%v\n", err)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Printf("read from resp.Body failed:%v", err)
// 		return
// 	}
// 	fmt.Println(string(body))

// }
