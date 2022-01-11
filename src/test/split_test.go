package test

/*在*_test.go文件中有三种类型的函数:
单元测试函数(函数名前缀为Test),   测试程序的一些逻辑行为是否正确
基准测试函数(函数名前缀为Benchmark),	测试函数的性能
示例函数(函数名前缀为Example),    为文档提供示例文档

//==============一、单元测试函数===========================
go test 测试输出
go test -v 测试输出详细信息
go test -v -run=Split/simple  只会运行simple对应的子测试用例。
go test -cover  查看测试覆盖率,你的代码被测试套件覆盖的百分比
go test -cover -coverprofile=c.out Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件
*/

import (
	"fmt"
	"reflect"
	"testing"
)

//单元测试：
func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":b")        // 程序输出的结果
	want := []string{"a", ":c"}        // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

//单元测试改进一：测试组
func TestChineseSplit(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%#v, got:%#v", tc.want, got)
		}
	}
}

//单元测试改进二：map测试
func TestMapSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

//单元测试改进三：子测试的Setup与Teardown
// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}
func TestOtherSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "枯藤老树昏鸦", sep: "老", want: []string{"枯藤", "树昏鸦"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("excepted:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

//==================二、基准测试函数============================
/*基准测试：
go test -bench=Split      执行基准测试
go test -bench=Split -benchmem 添加-benchmem参数，来获得内存分配的统计数据
*/

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}

/*性能比较函数： fib_test
上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，
比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？
再或者对于同一个任务究竟使用哪种算法性能最佳？
我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。
go test -bench=. 运行基准测试
默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。
BenchmarkFib40-12      2        516290250 ns/op  ns纳秒 op操作 ns/op，一次操作的时间。
1s=1*10^12ps
1s=1*10^9ns
1s=1*10^6us
1s=1*10^3ms
go test -bench=Fib40 -benchtime=20s
这一次BenchmarkFib40函数运行了50次，结果就会更准确一些了
*/

/*并行测试
func (b B) RunParallel(body func(PB))会以并行的方式执行给定的基准测试。
RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行，其中goroutine数量的默认值为GOMAXPROCS。
用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性，那么可以在RunParallel之前调用SetParallelism 。
RunParallel通常会与-cpu标志一同使用。
*/
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("枯藤老树昏鸦", "老")
		}
	})
}

//===========================三、实例函数=======================================
/*
示例函数能够作为文档直接使用，例如基于web的godoc中能把示例函数与对应的函数或包相关联。
示例函数只要有输出
例如包含了// Output:也是可以通过go test运行的可执行测试。
split $ go test -run Example
PASS
ok      github.com/pprof/studygo/code_demo/test_demo/split       0.006s
示例函数提供了可以直接运行的示例代码，可以直接在golang.org的godoc文档服务器上使用Go Playground运行示例代码。*/
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("枯藤老树昏鸦", "老"))
	// Output:
	// [a b c]
	// [枯藤 树昏鸦]
}

//==========================四、压力测试===========================
/*
go test -test.bench=".*"  表示测试全部的压力测试函数*/
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
