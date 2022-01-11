package go_learn

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*break 跳出本次循环，后面不执行
continue 跳出本次循环，后面执行
goto 跳出多层循环，到达指定位置
*/
func breaks() { //多重循环需要break几次才能跳出循环
	breakAgain := false
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if y == 2 {
				breakAgain = true
				break
			}
		}
		if breakAgain {
			break
		}
	}

}
func goto_func() { //使用goto一次就可以
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if y == 2 {
				goto breakHere
			}
		}
	}
	return
breakHere:
	fmt.Println("跳到这里")
}

func for_func() { //循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("for后接语句==>", sum)
}
func for_continued_func() { //for初始化语句何后置语句是可选的
	sum := 1
	for sum < 1000 { //2,4,8,16,32,64,128,256,..1024
		sum += sum
	}
	fmt.Println("for_continued_func,初始化语句何后置语句是可选的==>", sum)
}
func for_is_gos_while_func() { //C 的 while 在 Go 中叫做 for
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("sum==>", sum)
	}
	fmt.Println("while语句写法==>", sum)
}
func if_func(x float64) string { //if方法
	if x < 0 {
		return if_func(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func power(x, n, lim float64) float64 { //if后可接语句
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
func pow(x, n, lim float64) float64 { //if else方法
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g>=%g\n", v, lim)
	}
	return lim
}
func niudunfunc(x float64) float64 { //牛顿法求sqrt（），平方根
	z := 1.0
	j := 1
	for j <= 10 {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("第%d次计算的值为：%g\n", j, z)
		j += 1
	}
	return math.Sqrt(2)
}

/*switch语句后面接东西，case语句必须是确定值；
switch语句后啥都没有，case语句可以随意*/
func switch_func() { //switch的 case 语句从上到下顺次执行，直到匹配成功时停止
	fmt.Print("Go runs on  ")
	switch os := runtime.GOOS; os { //Go只运行选定的case，而非之后所有的case,
	case "darwin": //Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
		fmt.Println("OS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("没有合适的选项，输出在%s.\n", os)
	}
}
func switch_no() { //没有条件的 switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good night!")
	}
}
func defer_fun() {
	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	defer fmt.Println("world")
	fmt.Print("hello  ")
	fmt.Println("defer推迟的函数调用会被压入栈")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("出栈！")
}

func Res02() {
	for_func()
	for_continued_func()
	for_is_gos_while_func()
	fmt.Println("if语句==>", if_func(2), if_func(-4))
	fmt.Println(
		"if后可接简单语句==>",
		power(3, 2, 10),
		power(3, 3, 20),
	)
	fmt.Println(
		"在if的简短语句中声明的变量只能在对应的else块中使用:==>",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(niudunfunc(2))
	switch_func()
	switch_no()
	defer_fun()
}
