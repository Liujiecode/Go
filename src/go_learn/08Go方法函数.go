package go_learn

import (
	"fmt"
	"math"
	"reflect"
)

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

//s接收者，S表示s的类型
func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

func Res08() {
	/*值类型变量和指针类型变量都可以调用其值接收者方法和指针接收者方法，仅仅是Go语言的语法糖在起作用。
	  值接收者声明的方法，调用时会使用这个值的一个副本去执行，
	  而指针接收者在调用者会共享调用方法时接收者所指向的值，即可以修改指向的值。*/
	/*当实现了一个接收者是值类型的方法，就可以自动生成一个接收者是对应指针类型的方法，因为两者都不会影响接收者。
	  但是，当实现了一个接收者是指针类型的方法，如果此时自动生成一个接收者是值类型的方法，
	  原本期望对接收者的改变（通过指针实现），现在无法实现，因为值类型会产生一个拷贝，不会真正影响调用者。
	*/
	var s I = &S{}
	k := S{}
	fmt.Printf("%d %d\n", s.Get(), k.Get())
	s.Set(20)
	k.Set(30)
	fmt.Printf("%d %d\n", s.Get(), k.Get())
	// fmt.Println(s.Age)
}

/*匿名函数：匿名函数由一个不带函数名的函数声明和函数体组成
使用方法：1.匿名函数赋给一个变量，变量再调用
*/
func Res08_2() {
	//1.匿名函数赋给一个变量，变量再调用
	getsqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}
	fmt.Println(getsqrt(4), reflect.TypeOf(getsqrt(4)))
	//2.匿名函数自定调用，自己调用自己
	func() {
		for i, v := range []int{1, 2, 3, 4, 5} {
			fmt.Println(i, v)
		}
	}()

}

/*闭包，定义一个函数返回值就是自己内部定义的匿名函数
闭包：一个拥有许多变量和绑定了这些变量的环境的表达式（通常是一个函数），因而这些变量也是该表达式的一部分。
闭包=函数+引用环境
*/
func a2() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}
func a3() func() int {
	i := 0
	return func() int {
		i++
		fmt.Println(i)
		return i
	}
}
func Res08_03() {

	c := a2()
	c()
	c()
	c()
	d := a3()
	d()
	d()
	d()
}
