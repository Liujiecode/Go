package go_learn

import (
	"fmt"
)

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Scale3(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//带指针参数的函数，必须接收一个指针 Scale3Func(&v,5)
/*使用指针接收者的原因有二：
首先，方法能够修改其接收者指向的值。
其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。
通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。
*/
func Scale3Func(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Res09() {
	v := Vertex2{3, 4}
	v.Scale3(2)
	Scale3Func(&v, 10)

	p := &Vertex2{4, 3} //指针p表示指向Vertex2{4,3}的地址,*p指向Vertex2{4，3}的值
	p.Scale3(3)
	Scale3Func(p, 8)

	fmt.Println("v,*p,p:", v, *p, p)
}
