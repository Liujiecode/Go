package go_learn

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

/*映射将键映射到值。
映射的文法与结构体相似，不过必须有键名。
映射的零值为 nil 。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。*/

type Vertex1 struct {
	Lat, Long float32
}

//若顶级类型只是一个类型名，你可以在文法的元素中省略它。
/*练习：实现WordCount。它应当返回一个映射，其中包含字符串s中每个“单词”的个数。函数wc.Test会对此函数执行一系列测试用例，并输出成功还是失败。*/

func WordCount(s string) map[string]int {
	words := strings.Fields(s) //将字符串分割成多个子串
	wordcount := make(map[string]int)
	for _, word := range words {
		elem, ok := wordcount[word]
		//通过双赋值检测某个键是否存在：
		//elem, ok = p[key]
		//若 key 在 p 中，ok 为 true ；否则，ok 为 false。
		if ok == false {
			wordcount[word] = 1
		} else {
			wordcount[word] = elem + 1
		}
	}
	return wordcount
}

func Res05() {
	fmt.Println("---映射---")
	/*可以使用内建函数 make 也可以使用 map 关键字来定义 Map（集合）:
	var my_map map[string]int
	my_map:=make(map[string]int)
	如果不初始化map，那么就会创建一个nil map。nil map不能用来存放键值对；
	*/
	//写法一: var m map[string]Vertex1   m = make(map[string]Vertex1)
	var m map[string]Vertex1
	m = make(map[string]Vertex1)
	m["Bell Labs"] = Vertex1{
		40.684323, -74.39967,
	}
	m["test02"] = Vertex1{
		265465, 56.42,
	}
	//写法二
	var n = map[string]Vertex1{
		"test03": {40.111556, 40.34846465},
		"test04": {50.1, 50.3},
		"test05": {40, 52},
	}
	fmt.Println(m["Bell Labs"], m["test02"], m, n)
	//映射修改
	/*在映射 p 中插入或修改元素：
	  p[key] = elem
	  获取元素：
	  elem = p[key]
	  删除元素：
	  delete(p, key)
	  通过双赋值检测某个键是否存在：
	  elem, ok = p[key]
	  若 key 在 p 中，ok 为 true ；否则，ok 为 false。
	  若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	  同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	  注 ：若 elem 或 ok 还未声明，你可以使用短变量声明：
	  elem, ok := p[key]*/
	p := make(map[string]int)
	p["test06"] = 10
	fmt.Println("The Value:", p["test06"])
	p["test06"] = 100
	fmt.Println("Change The Value:", p["test06"])
	delete(p, "test06")
	fmt.Println("Delete The Value:", p["test06"])
	v, ok := p["test06"]
	fmt.Println("The Value:", v, "Present：", ok,
		"从映射读取不存在的值、值的类型：", p["01"], reflect.TypeOf(p["01"]), "映射:", p)
	list := "你好 hha 真巧 hha hh ok okde"
	fmt.Println(WordCount(list))

}
func Map() {
	/* 非并发下使用map
	   并发下使用sync.map      */
	//将键值对保存到sync.map
	var sc sync.Map
	sc.Store(01, "hello")
	sc.Store(02, "world")
	sc.Store(03, "!")
	sc.Store(04, "hehe")
	//从sync.map中根据键取值
	fmt.Println(sc.Load(01)) //hello
	//根据键删除对应的键值对
	sc.Delete(01)   //true
	fmt.Println(sc) //{{0 0} {{map[2:0xc000006090] false}} map[] 0}
	//遍历所有sc.map的键值对
	/*sync.Map没有提供获取map数量的方法，替代方法是获取时遍历自行计算数量*/
	var num int = 0
	sc.Range(func(k, v interface{}) bool {
		fmt.Println("sc: ", k, v)
		num += 1 //统计map数量
		return true

	})
	fmt.Println("num:", num)
}

/*反射reflect：
reflect.Type Of()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。*/
func Reflect() {
	var a int = 12
	/*反射的值对象（reflect.Value）*/
	//通过reflect.valueof()获取反射值对象
	Valueof_a := reflect.ValueOf(a)
	fmt.Println("Valueof_a: ", Valueof_a) //反射的值对象//12

	/*反射的类型对象（reflect.Type）*/
	Typeof_a := reflect.TypeOf(a)      //通过reflect.TypeOf()取得变量a的类型对象Typeof_a，类型为reflect.Type()
	fmt.Println("Typeof_a:", Typeof_a) //反射的类型对象 //int
	fmt.Printf("Typeof_a.Name():%v ,Typeof_a.Kind():%v\n", Typeof_a.Name(), Typeof_a.Kind())
	/*通过Typeof_a类型对象的成员函数，可以分别获取到Typeof_a变量的类型名为int，种类（Kind）为int。*/
	type Emm int
	type cat struct{}
	const Zero Emm = 0
	Typeof_Zero := reflect.TypeOf(Zero)
	Typeof_cat := reflect.TypeOf(cat{})

	fmt.Println(Typeof_Zero.Name(), Typeof_Zero.Kind(), Typeof_cat.Name(), Typeof_cat.Kind()) //Emm int cat struct

	/*通过反射修改变量值的前提条件之一：这个值必须可以被寻址。简单地说就是这个变量必须能被修改*/
	value_of_a := reflect.ValueOf(&a)
	value_of_a = value_of_a.Elem()
	value_of_a.SetInt(1)
	fmt.Println(value_of_a)
}
