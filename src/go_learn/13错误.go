package go_learn

import (
	"errors"
	"fmt"
	"math"
	"time"
)

/*实例一：*/
type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

/*实例二*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

/*
error是一个内建接口，自带的
type error interface{
	Error() string
}
*/
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func Res13() {
	fmt.Println("------实例一：")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println("-----实例二：")
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
func Panic() {
	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发 panic。
	panic(errors.New("something wrong")) //故意用panic引发异常，后面不会运行
	fmt.Println("Exit function main.")
}
func false01() {
	for {
		fmt.Println("1")
		panic("bug")     //主动抛出异常
		fmt.Println("4") //不会运行的.
		time.Sleep(1 * time.Second)
	}
}
func false02() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("xiaorui.cc start")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，"bug"
		}
		fmt.Println("xiaorui.cc end")
	}()
	for {
		fmt.Println("1")
		a := []string{"a", "b"}
		fmt.Println(a[3]) // 越界访问，肯定出现异常
		// panic("bug”)  // 上面已经出现异常了,所以肯定走不到这里了,下面代码无效。
		fmt.Println("4")
		time.Sleep(1 * time.Second)
	}
}

/*Panic使程序宕（dang）机，panic后面内容不再执行
Recover()使宕机恢复，继续下面的执行
*/
func Panic2() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("2")
		if err := recover(); err != nil {
			fmt.Println(err) //这里的err其实就是panic传入的内容，bug
		}
		fmt.Println("3")
	}()
	// false01()
	fmt.Println("--panic02--")
	false02()
}

/*统一异常处理*/
// func Errors_to_one(){
// 	err:=First Check Error()
// 	if err!=nil{
// 		goto on Exit
// 	}
// 	err=Second Check Error()
// 	if err!=nil{
// 		goto on Exit
// 	}
// 	fmt.print("Done")
// 	return
// on Exit:
// 	fmt.print("error")
// 	exit Process()
// }
