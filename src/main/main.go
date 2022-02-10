package main

import (
	"fmt"
	"go_learn"
)

/*一般测试*/
func main() {
	fmt.Println("主函数入口")
	// fmt.Println("-------")
	// str := "hello你好！"
	// fmt.Println([]byte(str))          //[104 101 108 108 111 228 189 160 229 165 189 239 188 129]
	// fmt.Println([]rune(str))          //[104 101 108 108 111 20320 22909 65281]
	// fmt.Println([]string{"hello你好！"}) //[hello你好！]
	// fmt.Println(strings.Fields(str))  //[hello你好！]
	// fmt.Println("-------")
	// go_learn.Res01_2()
	// str2 := "ni好啊"        //len(str2)==8
	// list1 := []byte(str2) //len(list1)==8
	// list2 := []rune(str2) //len(list2)==4
	// fmt.Println(len(list1), len(list2))
	// array01 := []int{1, 2, 3}
	// array02 := []int{4, 5, 6}
	// array01 = append(array01, array02...)
	// array03 := append(array01, array02[1:]...)
	// fmt.Println(array01) //[1 2 3 4 5 6]
	// fmt.Println(array03) //[1 2 3 4 5 6 5 6]
	// array04 := append(array03[:4], array03[5:]...)
	// fmt.Println(array04) //[1 2 3 4 6 5 6]
	// keys := make([]string, 0, 8)
	// fmt.Println(keys) //[]
	// keys2 := make([]string, 8)
	// fmt.Println(keys2) //[       ]
	// for key := range array03 {
	// 	key2 := strconv.Itoa(key)
	// 	keys = append(keys, key2)
	// }

	// sort.Strings(keys)
	// fmt.Println(keys) //[0 1 2 3 4 5 6 7]
	// fmt.Println("--slice--")
	// go_learn.Slice()
	// fmt.Println("--map--")
	// go_learn.Map()
	// // go_learn.Res00()
	// fmt.Println("--panic--")
	// go_learn.Panic()
	// go_learn.Panic2()
	// go_learn.Goto_func()
	// go_learn.Res08_2()
	// go_learn.Res08_03()
	go_learn.Atomic()
	go_learn.Res00()
	go_learn.Reflect()

}

// func testArrayPoint(x *[]int) {
// 	fmt.Printf("func Array:%p,%v\n", x, *x)
// 	(*x)[1] += 100
// }

// func main() {
/*
	go_learn.HHH()
	arrayA := []int{100, 200}
	testArrayPoint(&arrayA) // 1.传数组指针 func Array:0xc000096060,[100 200]
	arrayB := arrayA[:]
	fmt.Printf("arrayA:%p,%v", &arrayA, arrayA)       //arrayA:0xc000096060,[100 300]
	fmt.Printf("arratB:%p,%v\n", &arrayB, arrayB)     //arratB:0xc000096090,[100 300]
	testArrayPoint(&arrayB)                           // 2.传切片 func Array:0xc000096090,[100 300]
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA) //arrayA : 0xc000096060 , [100 400]

	var num int
	var p *int
	fmt.Println(num) //0
	p = &num
	*p = 30
	fmt.Println(num) //30

	//new()用法，x:=new(int) 定义了个*int的指针,并且初始化
	  //var x *int 定义了个*int的指针，没有初始化

	x := new(int)
	fmt.Println(*x, x) //0 0xc000016160

	var y *int
	fmt.Println(y) //<nil>
	var a map[string]int
	//a["he"] = 2      a只声明不初始化无法使用
	// var b map[string]int                //声明一个map类型的变量b，值nil
	b := make(map[string]int, 10) //对b初始化，值0，容量10
	b["a"] = 3
	fmt.Println(b)                          //map[a:3]
	c := map[string]int{"a": 1, "b": 2}     //声明并赋值，值{"a":1,"b":2,}，长度2
	fmt.Printf("a的值和大小：%v,%d\n", a, len(a)) //a的值和大小：map[],0
	fmt.Printf("b的值和大小：%v,%d\n", b, len(b)) //b的值和大小：map[a:3],1
	fmt.Printf("c的值和大小：%v,%d\n", c, len(c)) //c的值和大小：map[a:1 b:2],2
	// fmt.Println("=====map===")
	// usual.Map()
	// fmt.Println("===struct===")
	// usual.Struct()
	// fmt.Println("===defer===")
	// usual.Defer()
	fmt.Println("------------异常处理--------------")
	fmt.Println("--1.try catch异常处理--")
	// defer func() {
	// 	fmt.Println(recover())
	// }()
	// switch z, error := usual.Div(10, 0); error {
	// case nil:
	// 	fmt.Println(z)
	// case usual.ErrDivBuzero:
	// 	panic(error)
	// }
	// usual.Div_B()
	err := usual.Open("D:/MyCodeGroup/Go/src/usual/books.txt")
	switch v := err.(type) {
	case *usual.PathError:
		fmt.Println("get path error,", v)
	default:
		fmt.Println("路径正确！")

	}
	fmt.Println("------------单元测试--------------")
*/
// fmt.Println("--------------接口----------------")
// var x usual.Sayer
// a := usual.Cat{}
// b := usual.Dog{}
// x = a
// fmt.Print("猫🐱 ：")
// x.Say()
// x = b
// fmt.Print("狗🐕 ：")
// b.Say()
// usual.Type_assertion()
// fmt.Println("--------------服务端编程----------------")
// resp, _ := http.Get("http://www.baidu.com")
// //fmt.Println(resp)
// // resp, _ := http.Get("http://127.0.0.1:8000/go")
// defer resp.Body.Close()
// // 200 OK
// fmt.Println(resp.Status)
// fmt.Println(resp.Header)
// fmt.Println(resp.Request.URL)

// buf := make([]byte, 1024)
// for {
// 	// 接收服务端信息
// 	n, err := resp.Body.Read(buf)
// 	if err != nil && err != io.EOF {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		fmt.Println("读取完毕")
// 		res := string(buf[:n])
// 		fmt.Println(res)
// 		break
// 	}
// }
// go_learn.Main_Goruntine()
// go_learn.Main_channel()
// go_learn.Add_test()
// var s string
// var s1 byte
// var s2 int
// s = "2"
// s1 = 'a' //int('a')=97
// s2 = 2

// fmt.Printf("s--%T s1--%T s2--%T\n", s, s1, s2)
// fmt.Printf("s--%T int(s1)--%v\n", s, int(s1))
// fmt.Printf("a--%v 2--%v strconv--%v \n", int('a'), int('2'), strconv.Itoa(12))
// b := strconv.Itoa(12)
// fmt.Printf("%v, %s, %T\n", b[0], b, b[0])
// var s3 string = "golang"
// s4 := []byte(s3)
// s5 := string(s4)

// fmt.Printf("s3--%v s4--%v s5--%v ", s3, s4, s5)
// go_learn.Main_Time()
// go_learn.Main_Select()
// go_learn.Main_Lock()
// go_learn.Main_Sync()
// go_learn.Main_Sync_Map()
// go_learn.Res08_2()
// go_learn.Res08_03()
// usual.Main_Robot()
// usual.GetEmail()
// mysql.Main_mysql01()

// }
