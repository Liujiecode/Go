package go_learn

import (
	"fmt"
	"reflect"
	"strconv" //类型转换
	"strings"
)

type X interface{}

func test01(id int) X {
	strconv.Itoa(id)
	return id
}
func test() {
	// strconv.Atoi(id)
	strconv.Atoi("nisad")         //string-->int
	strconv.CanBackquote("nisad") //string-->bool
	strconv.FormatBool(true)      //bool-->string
	strconv.Itoa(12)              //int-->string

}

type IPAddr [4]byte

func Res00() {
	fmt.Println("Hello, World!")
	/*Go语言数据类型*/
	// bool

	// string

	// int		有符号整型
	// int8		有符号8位整型
	// int16
	// int32
	// int64

	// uint		无符号整型
	// uint8	无符号8位整型(0 到 255)
	// uint16
	// uint32
	// uint64

	// uintptr   无符号整型，用于存放一个指针

	// byte    类似 uint8

	// float32		32位浮点型数
	// float64

	// complex64	32位实数和虚数
	// complex128	 64位实数和虚数
	/*查看类型*/
	var m = 15
	//法一：
	fmt.Printf("m的类型为：%T\n", m)
	//法二：import "reflect"
	fmt.Println("m的类型为：", reflect.TypeOf(m))
	/*strings.Fields()*/
	s := "hello world hello world"
	//str := "wo"
	//以连续的空白字符为分隔符，将s切分成多个子串，结果中不包含空白字符本身。
	//空白字符有：\t, \n, \v, \f, \r, ’ ‘, U+0085 (NEL), U+00A0 (NBSP) 。
	//如果 s 中只包含空白字符，则返回一个空列表
	index := strings.Fields(s)
	fmt.Println(index, len(index)) //4
	host := map[string]IPAddr{
		"LOOKBACK": {127, 0, 0, 1},
		"DNS":      {8, 8, 8, 8},
	}
	for name, ip := range host {
		fmt.Printf("%v:%v\n", name, ip)
	}

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.Atoi("s"))

}

/*  将 IPAddr{1，2，3，4}转成 "1.2.3.4"    */
func (i IPAddr) String() string {
	var s string
	for idx, val := range i {
		if idx != 3 {
			s += strconv.Itoa(int(val)) + "."
		} else {
			s += strconv.Itoa(int(val))
		}
	}
	return s
}
