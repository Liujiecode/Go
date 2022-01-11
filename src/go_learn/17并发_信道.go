package go_learn

import (
	"fmt"
)

/*
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）
和映射与切片一样，信道在使用前必须创建：
ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得Go程可以在没有显式的锁或竞态变量的情况下进行同步。
以下示例对切片中的数进行求和，将任务分配给两个Go程。一旦两个Go程完成了它们的计算，它就能算出最终的结果。
*/
/*
信道可以是带缓冲的。将缓冲长度作为第二个参数提供给make来初始化一个带缓冲的信道：
ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

/*
range 和 close
发送者可通过close关闭一个信道来表示没有需要发送的值了。
接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完
v, ok := <-ch
之后ok会被设置为 false。
循环for i := range c 会不断从信道接收值，直到它被关闭。
*注意：只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。
*注意：信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个range循环。*/
func fibonacci(n int, d chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		d <- x
		x, y = y, x+y
	}
	close(d)
}
func Res17() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //17
	go sum(s[len(s)/2:], c) //-5
	x, y := <-c, <-c        // 从 c 中接收
	fmt.Println(x, y, x+y)
	fmt.Println("---带缓冲的信道---")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 //也会填满缓冲区
	//fmt.Println("三个信道长度，填满缓冲区",<-ch, <-ch, <-ch)  //fatal error: all goroutines are asleep - deadlock!
	fmt.Println("输出信道ch", ch)
	fmt.Println("两个信道长度，没有填满缓冲区", <-ch, <-ch)
	fmt.Println("---信道关闭---")
	d := make(chan int, 10)
	fmt.Println("cap方法：", cap(d)) //cap（）计算长度
	go fibonacci(cap(d), d)
	//fmt.Println("直接输出d", d, "输出<-d", <-d, "输出<-d，again", <-d, <-d)
	for i := range d { //通过range遍历信道
		fmt.Print(i, " ")
	}

	fmt.Println("d,<-d", d, <-d, <-d, <-d, <-d)

}
