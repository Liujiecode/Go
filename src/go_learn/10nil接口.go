package go_learn

import "fmt"

type I2 interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I2，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}
func (t *T) N() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func dedcribe(i I2) {
	fmt.Printf("(%v,%T)\n", i, i)
}

func Res10() {
	var i I2 = T{"hello"}
	c := T{"world"}
	i.M()
	c.M()

	var j I2
	var t *T
	j = t       //值为空，底层值为 nil 的接口值
	dedcribe(j) //(<nil>, *main.T)
	j.M()       //<nil>
	j = &T{"Hello"}
	dedcribe(j) //(&{hello}, *main.T)
	j.M()       //hello

}
