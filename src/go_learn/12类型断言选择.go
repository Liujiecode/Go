package go_learn

import "fmt"

/*类型断言 提供了访问接口值底层具体值的方式。

t := i.(T)
该语句断言接口值i保存了具体类型T，并将其底层类型为T的值赋予变量t。
若i并未保存T类型的值，该语句就会触发一个恐慌。

为了判断一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。
t, ok := i.(T)  两个参数 t：底层值/0，ok：true/false
若i保存了一个T，那么t将会是其底层值，而ok为true。

否则，ok将为false而t将为T类型的零值，程序并不会产生恐慌。
请注意这种语法和读取一个映射时的相同之处。*/
func do(i interface{}) {
	switch v := i.(type) { //type关键字
	case int:
		fmt.Printf("两次 %v 是%v\n", v, v*2)
	case string:
		fmt.Printf("%q 是 %v bytes long\n", v, len(v))
	default:
		fmt.Printf("我不知道类型%T\n", v)
	}

}

func Res12() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) //hellp

	s, ok := i.(string)
	fmt.Println(s, ok) //hello true

	f, ok := i.(float64) //类型不同，接口i未保存具体类型float64
	fmt.Println(f, ok)   //0 false

	// f = i.(float64) // 报错(panic)
	// fmt.Println(f)
	do(20)
	do("hello")
	do(true)

}
