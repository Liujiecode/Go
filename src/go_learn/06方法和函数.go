package go_learn

import (
	"fmt"
	"math"
)

/*Go没有类。不过你可以为结构体类型定义方法。
方法就是一类带特殊的接收者参数的函数。
方法接收者在它自己的参数列表内，位于func关键字和方法名之间。
在此例中，Abs方法拥有一个名为v，类型为Vertex的接收者。
//结构体：type 结构体名 struct{}
//方法：func (接收者) 方法名 返回类型{}
*/
type okk struct {
	a, b, c int
}

type Vertex struct {
	X, Y float64
}

//方法：func (接收者) 方法名 返回类型{}
func (v Vertex) Abs() float64 {
	// fmt.Println("算术平方根求解：sqrt（16）",math.Sqrt(16)) //math.sqrt(m),求m的算术平方根
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//指针接收者的方法可以修改接收者指向的值（就像 Scale 在这做的）。
//由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
func (v *Vertex) Scale(f float64) { //加*,执行scale后会改变接收者指向的值，不加*不会改变
	v.X = v.X * f
	v.Y = v.Y * f
}
func (v Vertex) Scale2(f float64) { //加*,执行scale后会改变接收者指向的值，不加*不会改变
	v.X = v.X * f
	v.Y = v.Y * f
}
func Res06_1() {
	fmt.Println("算术平方根求解：sqrt（16）", math.Sqrt(16)) //math.sqrt(m),求m的算术平方根
	v := Vertex{3, 4}
	o := okk{2, 2, 2}
	v.Scale2(10)
	fmt.Println("求v:=Vertex{3，4}执行Scale2（10）后的Abs（）值", v.Abs())
	fmt.Println("输出o,v", o, v)
	v.Scale(10)
	fmt.Println("求v：=Vertex{3，4}，执行Scale（10）后的Abs（）的值", v.Abs())
}
