package go_learn

import "fmt"

//定义结构体
type People struct {
	name string
	age  int
}

//定义方法
func (p People) my_method() (string, int) {
	return "p的名字为：" + p.name + "  p的年纪为：", p.age
}

//定义函数
func my_function(n People) string {
	return "p的名字为:" + n.name
}
func Res06_2() {
	var p People
	p.age = 16
	p.name = "GoLang"
	fmt.Println(p.my_method())  //方法必须.出来
	fmt.Println(my_function(p)) //函数可以直接执行

}
