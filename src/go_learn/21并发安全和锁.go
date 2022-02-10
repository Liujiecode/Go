package go_learn

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var x int64

/*常用的控制共享资源访问的方法：
1.原子访问（atomic包）
2.互斥锁（sync.Mutex）
3.等待组（sync.Wait Group）
*/
var wg sync.WaitGroup //定义等待组，实现并发任务的同步（计时器）
/*互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
Go语言中使用sync包的Mutex类型来实现互斥锁。
使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁；
当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。*/
var lock sync.Mutex     //定义互斥锁
var rwlock sync.RWMutex //定义读写锁，非常适合读多写少的场景
//原子操作
var seq int64

func Gen_ID() int64 { //生成10个并发序列号
	// res := atomic.AddInt64(&seq, 1)
	// return res  //1 3 5 7 9 11 13 15 17 19
	return atomic.AddInt64(&seq, 1)

}
func Atomic() {
	for i := 0; i < 10; i++ {
		go Gen_ID()
		fmt.Print(Gen_ID(), " ")

	}

}

func add2() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}
func write2() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read2() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}
func Main_Lock() {
	wg.Add(2)
	go add2() //两个协程同时执行add2（）方法
	go add2()
	wg.Wait()
	fmt.Println(x) //10000 不加锁会导致输出老在变，同一公共资源回抢占
	/*读写互斥锁*/
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write2()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read2()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}

/*使用sync.WaitGroup来实现并发任务的同步:
方法名	                功能
(wg * WaitGroup) Add(delta int)	计数器+delta
(wg *WaitGroup) Done()	计数器-1
(wg *WaitGroup) Wait()	阻塞直到计数器变为0
sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N 个并发任务时，就将计数器值增加N。
每个任务完成时通过调用Done()方法将计数器减1。
通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。
*/
func hello2() {
	defer wg.Done()
	fmt.Println("Hello Goroutine!")
}
func Main_Sync() {
	wg.Add(1)
	go hello2() //// 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg.Wait()

}

/*sync.Once加载文件配置：防止初始化操作被加载多次
sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。
这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
var icons map[string]image.Image
var loadIconsOnce sync.Once

func loadIcons() {
    icons = map[string]image.Image{
        "left":  loadIcon("left.png"),
        "up":    loadIcon("up.png"),
        "right": loadIcon("right.png"),
        "down":  loadIcon("down.png"),
    }
}
// Icon 是并发安全的
func Icon(name string) image.Image {
	// if icons == nil {  这种处理就容易出错
    //     loadIcons()
    // }
    // return icons[name]
	//sync.Once其实内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。
	//这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
*/
/*sync.Map
Go语言中内置的map不是并发安全的,少量几个goroutine的时候可能没什么问题，当并发多了之后执行上面的代码就会报fatal error: concurrent map writes错误。
*/
var map1 = make(map[string]int)
var map2 = sync.Map{}

func get(key string) int {
	return map1[key]
}

func set(key string, value int) {
	map1[key] = value
}

func Main_Sync_Map() {
	// wg := sync.WaitGroup{}
	// for i := 0; i < 20; i++ {
	// 	wg.Add(1)
	// 	go func(n int) {
	// 		key := strconv.Itoa(n)
	// 		set(key, n)
	// 		fmt.Printf("k=:%v,v:=%v\n", key, get(key))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			map2.Store(key, n)
			value, _ := map2.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i) //匿名函数自调用
	}
	wg.Wait()
}
