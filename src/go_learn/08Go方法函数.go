package go_learn

import "fmt"

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
