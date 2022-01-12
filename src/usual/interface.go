package usual

import "fmt"

/*接口（interface）是一种类型，一种抽象的类型。
接口就是一个需要实现的方法列表。
值接收者和指针接收者实现接口的区别：
使用值接收者实现接口之后，结构体类型变量 和 结构体指针类型变量 都可以赋值给该接口变量；
使用指针接收者实现接口之后，只有 结构体指针类型的变量 可以赋值给该接口变量。
一个类型可以定义多个接口，多个类型也可以实现同一个接口。
只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。
*/
//定义空接口：可以存储任意类型值
var null_interface interface{}

/*1.空接口作为函数的参数;
2.使用空接口实现可以保存任意值的字典。
3.类型断言：判断空接口中的值的类型与猜的是否一致。
*/
func Type_assertion() {
	var x interface{} = "hello go"
	// x = "hello go"
	v, ok := x.(string) //ok是判断x是不是string类型
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败，类型判断错误！")
	}
}

type Sayer interface {
	Say()
}
type Dog struct{}
type Cat struct{}

func (d Dog) Say() {
	fmt.Println("汪汪汪！")
}
func (c Cat) Say() {
	fmt.Println("喵喵喵！")
}
