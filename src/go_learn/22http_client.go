package go_learn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*client(客户端)*/
func Client01() {
	resp, err := http.Get("https://127.0.0.1:9090/client/")
	/*对于Get请求，参数都放在url上*/
	if err != nil {
		fmt.Printf("get url failed,%v\n", err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body) //客户端读出服务端响应的body
	if err != nil {
		fmt.Printf("read resp.Body failed,%v\n", err)
		return
	}
	fmt.Println(string(data))
}
