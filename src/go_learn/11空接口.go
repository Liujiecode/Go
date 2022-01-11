package go_learn

import "fmt"

func Res11() {
	//指定了零个方法的接口值被称为 *空接口：*
	/*空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）
	空接口被用来处理未知类型的值。例如，fmt.Print可接受类型为interface{}的任意数量的参数。*/
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
