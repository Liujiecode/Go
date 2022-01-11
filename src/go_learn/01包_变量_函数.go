package go_learn

import (
	"fmt"
	"math"
)

/*位运算
<<n 表示左移n位  101<<1 ==110 101<<2=1100  101<<3=11000
*/

/*常量*/
const Pi = 3.14 //声明常量
const width, height, bo, str = 10, 20, true, "你好！"

/*逻辑运算符
&& 逻辑and
|| 逻辑or
！ 逻辑not
*/
/*位运算符
& 与运算，全真为真
| 或运算，全假为假
^ 异或运算，相同为假，不同为真
假定 A=60, A=0011 1100
   B=13, B=0000 1101
   A&B =0000 1100
   A|B =0011 1101
   A^B =0011 0001
   ~A =1100 0011
   A<<2 =1111 0000  左移，移动最右边的1两次，移动一次计算一次0111 1000==》1111 0000或全部左移两位，高位丢弃，低位补0
   A>>2 =0000 1111  右移，移动最左边的1两次，移动一次计算一次0001 1110==》0000 1111或全部右移两位
*/

/*特殊常量iota*/
//第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
const (
	a = iota //0
	b        //1
	c        //2
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
) //0 1 2 ha ha 100 100 7 8

func add(x int, y int) int { //func add(x,y int)int{ return x+y}
	return x + y
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

/*函数的闭包*/
/*Go函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*闭包函数实例：斐波那契闭包*/
func fibaonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a
	}

}

func Res01_1() {
	//var a, b, c int
	var m int //var语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
	var n bool
	var k int = 5 //k:=5 声明变量
	p := 5
	fmt.Println(math.Pi)
	fmt.Println(add(41, 20))
	fmt.Println(m, n, k, p)
	fmt.Println("---函数当值传递---")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))      //13
	fmt.Println(compute(hypot))    //5
	fmt.Println(compute(math.Pow)) //81
	fmt.Println("---函数的闭包---")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	fmt.Println("--闭包实例：斐波那契实例(0,1,1,2,3,5...)--")
	f := fibaonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
