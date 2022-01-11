package go_learn

import "fmt"

/*指针*/
//指针保存了值的内存地址
func pointer() {
	i, j := 42, 2701

	p := &i            // 指针p指向i的地址，*p表示i的值
	fmt.Print(*p, " ") // 通过指针读取 i 的值
	*p = 21            // 通过指针设置 i 的值
	fmt.Print(i, " ")  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

/*结构体*/
//一个结构体（struct）就是一组字段（field）。
/*结构体指针*/
// 结构体字段可以通过结构体指针来访问。
// 如果我们有一个指向结构体的指针p，那么可以通过(*p).X来访问其字段X。
// 不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写p.X就可以
/*结构体文法*/
//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合
type Verter struct {
	X int
	Y int
}

var ( //结构体表达方式
	v1 = Verter{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Verter{X: 1}  // Y:0 被隐式地赋予
	v3 = Verter{}      // X:0 Y:0
	p  = &Verter{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	m  Verter
)

func Res03() {
	pointer()
	fmt.Println(Verter{1, 2})
	v := Verter{3, 4}
	v.X = 6 //结构体字段使用点号来访问。
	fmt.Println("结构体字段使用点号访问", v.X)
	p := &v
	p.X = 12
	fmt.Println("结构体指针", v)
	fmt.Println(v1, p, v2, v3)
	//range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		//当使用for循环遍历切片时，每次迭代都会返回两个值。
		//第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
		//可以将下标或值赋予 _ 来忽略它。
		// for i, _ := range pow
		// for _, value := range pow
		fmt.Printf("2**%d = %d\n", i, v)
	}
	//可以将下标或值赋予 _ 来忽略它。
	// for _, v := range pow {
	// 	fmt.Printf("2**%d=%d\n", _, v)
	// }

}
