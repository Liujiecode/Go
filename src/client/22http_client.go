package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*使用一个公有的client，实用于请求很频繁，比如5s一次*/
// var (
// 	client = http.Client{
// 		Transport: &http.Transport{
// 			DisableKeepAlives: false,
// 		},
// 	}
// )

/*client(客户端)*/
func main() {
	/*发送请求：方式一*/
	/*对于Get请求，参数都放在url上*/
	// resp, err := http.Get("http://127.0.0.1:9090/client/?name=sb&age=18")
	// if err != nil {
	// 	fmt.Printf("get url failed,%v\n", err)
	// 	return
	// }
	/*发送请求：方式二，当参数带有？=咋办*/
	urlobj, err := url.Parse("http://127.0.0.1:9090/client/")
	data := url.Values{}
	data.Add("name", "超人？")
	data.Add("age", "?=100?")
	queryStr := data.Encode()
	fmt.Println("queryStr", queryStr)
	// fmt.Printf("urlobj.RawPath :%v, urlobj.RawQuery: %v\n", urlobj.RawPath, urlobj.RawQuery)
	fmt.Println("urlobg:", urlobj)
	urlobj.RawQuery = queryStr
	fmt.Println("urlobj_最终", urlobj)
	req, err := http.NewRequest("GET", urlobj.String(), nil)
	if err != nil {
		fmt.Println("http.newrequest wrong")
		return
	}
	//----------------常规写法---------------------
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Println("http.DefaultClient.Do(req) wrong")
	// 	return
	// }
	//-------------------------------------------------
	/*适用于请求不是很频繁，用完就关闭，比如一天两次请求*/
	/*禁用Keepalives的client*/
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response failed!")
		return
	}
	//-----------------------------------------------

	defer resp.Body.Close()
	datas, err := ioutil.ReadAll(resp.Body) //客户端读出服务端响应的body

	if err != nil {
		fmt.Printf("read resp.Body failed,%v\n", err)
		return
	}
	fmt.Println(string(datas))
}
