package go_learn

import (
	"fmt"
	maths "math/rand"
	"runtime"
	"time"
)

/* 进程 > 线程 > 协程
goroutine 奉行通过通信来共享内存，而不是共享内存来通信。
一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数。
*/
func hello() {
	fmt.Println("Hello GoRoutine!")
}
func a1() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b1() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

/*
在程序启动时，Go程序就会为main()函数创建一个默认的goroutine。
当main()函数返回的时候该goroutine就结束了，所有在main()函数中启动的goroutine会一同结束。
*/
func Main_Goruntine() { //在程序启动时，Go程序就会为main()函数创建一个默认的goroutine

	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	time.Sleep(time.Second) //先打印main goroutine done!是因为我们在创建新的goroutine的时候需要花费一些时间，而此时main函数所在的goroutine是继续执行
	fmt.Println("----runtime.Gosched()----")
	/*让出CPU时间片，重新等待安排任务(大概意思就是本来计划的好好的周末出去烧烤，但是你妈让你去相亲,
	两种情况第一就是你相亲速度非常快，见面就黄不耽误你继续烧烤，
	第二种情况就是你相亲速度特别慢，见面就是你侬我侬的，耽误了烧烤，但是还馋就是耽误了烧烤你还得去烧烤)*/
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
	fmt.Println("----runtime.GoMaxProcs----")
	cpu_nums := runtime.NumCPU()
	fmt.Println("查询cpu数目：", cpu_nums)
	/*Go语言中可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数。*/
	runtime.GOMAXPROCS(2)
	go a1()
	go b1()
	time.Sleep(time.Second)
	fmt.Println("----runtime.Goexit()----")
	/*退出当前协程(一边烧烤一边相亲，突然发现相亲对象太丑影响烧烤，果断让她滚蛋，然后也就没有然后了)*/
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功！", ret)
}
func Main_channel() {
	/*无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。
	使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
		ch := make(chan int)
		// 创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值。
		// 就像你住的小区没有快递柜和代收点，快递员给你打电话必须要把这个物品送到你的手中，
		// 简单来说就是无缓冲的通道必须有接收才能发送。
		ch <- 10 //代码会阻塞在ch <- 10这一行代码形成死锁
		fmt.Println("发送成功！")*/
	fmt.Println("----无缓冲通道：有发送必须有接收----")
	ch := make(chan int)
	go recv(ch) //启用goroutine从通道接受值
	ch <- 10
	fmt.Println("发送成功！")
	fmt.Println("----有缓冲通道----")
	/*我们可以在使用make函数初始化通道的时候为其指定通道的容量*/
	ch2 := make(chan int, 1)
	ch2 <- 10
	fmt.Println("发送成功！")
	fmt.Println("----close()----")
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")
	fmt.Println("----通道取值练习----")
	ch3 := make(chan int)
	ch4 := make(chan int)
	go func() {
		for i := range []int{1, 2, 3, 4, 5, 6, 7} {
			ch3 <- i
		}
		close(ch3)
	}()
	go func() {
		for {
			i, ok := <-ch3
			if !ok {
				break
			}
			ch4 <- i * i
		}
		close(ch4)
	}()
	for i := range ch4 {
		fmt.Println(i)
	}
	fmt.Println("----单向通道----")
	ch5 := make(chan int)
	ch6 := make(chan int)
	go counter(ch5)
	go squarer(ch6, ch5)
	printer(ch6)

}
func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
func Add_test() {
	fmt.Println("----Goroutine池----")
	jobChan := make(chan *Job, 5)
	// 2.结果管道
	resultChan := make(chan *Result, 5)
	// 3.创建工作池
	createPool(2, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := maths.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}

}
