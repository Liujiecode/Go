package go_learn

import (
	"fmt"
	"strings"
)

/*切片的长度与容量
切片的长度就是它所包含的元素个数。
切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
切片s的长度和容量可通过表达式 len(s) 和 cap(s) 来获取；
切片追加元素
var s []int  //[]
append(s,0,1)  //[0,1]*/

func slice_len_cap() {
	print_len_cap := func(s []int) { //函数作为变量
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
	s := []int{2, 3, 5, 7, 11, 13}
	print_len_cap(s)

	// 截取切片使其长度为 0
	s = s[:0]
	print_len_cap(s)

	// 拓展其长度
	s = s[:4]
	print_len_cap(s)

	// 舍弃前两个值
	s = s[2:]
	print_len_cap(s)

}

/*make创建切片,make([]type,len,cap)*/
func making_slice() {
	fmt.Println("----making创建切片----")
	print_len_cap := func(s string, x []int) { //函数作为变量
		fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
	}
	a := make([]int, 5)
	print_len_cap("a", a)

	b := make([]int, 0, 5)
	print_len_cap("b", b) //b是空数组

	c := b[:2] //给c两个值的长度
	print_len_cap("c", c)

	d := c[2:5] //给d三个值的长度
	print_len_cap("d", d)
}

/*切片的切片*/
func slice_to_slice() {
	board := [][]string{
		{"_", "_", "_"}, //[]string{"_", "_", "_"},无需再定义类型
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " ")) //将每一行拼接空字符成为新的字符串
		fmt.Println(board[i])                           //一行列表
	}
}

func Res04() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Go"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 9, 11}
	fmt.Println(primes) //结果：[2 3 5 7 9 11]
	/*切片a[low:high],前闭后开区间*/
	/*切片并不存储任何数据，它只是描述了底层数组中的一段。
	更改切片的元素会修改其底层数组中对应的元素。
	与它共享底层数组的切片都会观测到这些修改*/
	var s []int = primes[1:4] //表示[1:4)
	fmt.Println(s)            //结果：[3 5 7]
	fmt.Println("----更改切片元素会修改底层数组元素----")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) //[John Paul George Ringo]

	c := names[0:2]   //[John Paul]
	d := names[1:3]   //[Paul George]
	fmt.Println(c, d) //[John Paul] [Paul George]

	d[0] = "XXX"
	fmt.Println(a, d)  //[Hello Go] [XXX George]
	fmt.Println(names) //[John XXX George Ringo]
	fmt.Println("---切片的长度与容量----")
	slice_len_cap()
	/*切片的零值==nil*/
	//null_list := []int{} //赋值一个空切片/空数组
	var null_list []int //定义一个空数组/空切片
	if null_list == nil {
		fmt.Println("切片为空时==nil，类似于true和false")
	}
	making_slice()
	fmt.Println("---切片的切片---")
	slice_to_slice()
}
func Slice() {
	slice01 := []int{}
	for i := 0; i < 5; i++ {
		slice01 = append(slice01, i) //切片添加元素
	}
	fmt.Println(slice01)
	/*Go语言中切片删除元素的本质是：以被删除元素为分界点，将前后两个部分的内存重新连接起来。*/ //[0 1 2 3 4]
	slice01 = append(slice01[:1], slice01[2:]...)    //切片删除指定元素
	fmt.Println(slice01)

}
