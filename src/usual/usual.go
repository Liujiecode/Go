package usual

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/*Map是一种通过key来获取value的一个数据结构，其底层存储方式为数组，
在存储时key不能重复，当key重复时，value进行覆盖*/
func Map() {
	scoremap := make(map[string]int)
	scoremap["张三"] = 60
	scoremap["李四"] = 70
	scoremap["王五"] = 80
	//1.遍历map
	for k, v := range scoremap {
		fmt.Print(k, v, " ")
	}
	for k := range scoremap {
		fmt.Print(k, " ")
	}
	for _, v := range scoremap {
		fmt.Print(v, " ")
	}
	//2.判断某个键是否存在
	v, ok := scoremap["张三"]
	if ok {
		fmt.Println("存在", v)
	} else {
		fmt.Println("不存在！")
	}
	//3.删除
	delete(scoremap, "张三")
	//4.map类型的切片
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	//5.值为切片类型的map
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

type student struct {
	name string
	age  int
}
type person struct {
	name string
	city string
	age  int8
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

/*结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。*/
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}
type Student2 struct {
	ID     int
	Gender string
	Name   string
}
type Class struct {
	Title    string
	Students []*Student
}
type Class2 struct {
	Title    string
	Students []*Student2
}

func Struct() {
	//结构体定义
	type people struct {
		name string
		city string
		age  int8
	}
	//结构实例化
	var p1 people
	p1.name = "张三"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v p1=%+v p1=%#v\n", p1, p1, p1)
	//输出：p1={张三 北京 18} p1={name:张三 city:北京 age:18} p1=usual.people{name:"张三", city:"北京", age:18}
	/*1.%v当碰到数组时，仅输出value，不输出key;
	2.%+v当碰到数组时，将key-value 都输出;
	3.%#v输出时，会将方法名+k/v都输出*/
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	for _, stu := range stus {
		fmt.Print("stu==> ", stu)
		fmt.Println("&stu==>", &stu)
		m[stu.name] = &stu
	}
	fmt.Println("m:  ", m, &m)
	for k, v := range m {
		fmt.Println("v的指针==>", &v)
		fmt.Println(k, "=>", v.name, v.age)
	}
	p9 := newPerson("pprof.cn", "测试", 90)
	fmt.Printf("%#v\n", p9)
	//结构体与json序列化
	s1 := Student{
		ID:     1,
		Gender: "女",
		name:   "Tom",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)

	c := &Class2{
		Title:    "101",
		Students: make([]*Student2, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu2 := &Student2{
			Name:   fmt.Sprintf("stu2%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu2)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data2, err2 := json.Marshal(c)
	if err2 != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data2)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}

/*defer 延迟调用*/
func Defer() error {
	f, err := os.Open("books.txt")
	fmt.Println(f)
	if err != nil {
		return err
	}
	if f != nil {

		fmt.Println(f)
		defer f.Close()
	}
	return nil
}

/*异常处理error  或  panic-recover
Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。
如何区别使用 panic 和 error 两种方式?
惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。*/
/*标准库errors.New和fmt.Errorf函数用于创建实现error接口的错误对象。
通过判断错误对象实例来确定具体错误类型。*/
// var ErrDivBuzero = errors.New("Division by 0")
var ErrDivBuzero = fmt.Errorf("你错了！！")

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivBuzero
	}
	return x / y, nil
}
func getCircleArea(radius float32) (area float32) {
	if radius < 0 { // 自己抛
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}
func Div_B() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	getCircleArea(-5)
	fmt.Println("这里有没有执行")
}

//自定义error
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}
