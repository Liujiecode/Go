package go_learn

import (
	"fmt"
	"reflect"
	"strings"
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
