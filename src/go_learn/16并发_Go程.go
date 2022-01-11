package go_learn

import (
	"fmt"
	"time"
)

/*
Go 程（goroutine）是由Go运行时管理的轻量级线程
Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
sync 包提供了这种能力，不过在Go中并不经常用到，因为还有其它的办法*/
/*
select 语句使一个 Go 程可以等待多个通信操作。
select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
当 select 中的其它分支都没有准备好时，default分支就会执行。
为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：*/
func fibonacci02(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			// default:
			// 	fmt.Println("都没执行到！")
			// 	return
		default:
			fmt.Print("-->")
			time.Sleep(50 * time.Millisecond)

		}
	}
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func Res16() {
	go say("world") //输出的 hello 和 world 是没有固定先后顺序的
	say("hello")
	c := make(chan int) //创建通道（信道）
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c, " ")
		}
		quit <- 0
	}()
	fibonacci02(c, quit)
	fmt.Println("50*time.Millisecond:", 50*time.Millisecond,
		"\n50*time.Microsecond:", 50*time.Microsecond,
		"\n/50*time.January:", 50*time.January,
		"\ntime.January:", time.January,
		"\n50*time.Hour:", 50*time.Hour,
		"\ntime.Hour:", time.Hour)
}
