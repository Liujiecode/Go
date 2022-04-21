package go_learn

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Str_Test() {
	/*fmt.Sprintf  fmt.Sprint  fmt.Sprintln 返回字符串*/
	//go 中格式化字符串并赋值给新串，使用 fmt.Sprintf()
	var Name = "Tom"
	var Age = 18
	var url = "小白%s和小黑%d"
	res := fmt.Sprintf("%s is %d years old", Name, Age)
	res2 := fmt.Sprintf(url, Name, Age)
	//字符串拼接，当相邻两个操作数都不是字符串时，在操作数之间添加空格。
	res3 := fmt.Sprint(15, 12, 6, 7)
	res4 := fmt.Sprint("ss", 15, 12, 6, 7)
	res5 := fmt.Sprint("aa", 15, "a", 16, 17)
	//在操作数之间总是添加空格，输出添加换行符。
	res6 := fmt.Sprintln("asd", 15, "asd", 16, 17)
	res7 := fmt.Sprintln("hhhh", "ssss")
	fmt.Print("fmt.Sprintf=> ", res, "=> ", res2, "\n",
		"fmt.Sprint=> ", res3, "=> ", res4, "=> ", res5, "\n",
		"fmt.Sprintln=> ", res6, res7)

}
func Res01_2() {
	s1 := "hello world" //单行符串
	s2 := `ni
	hao
	go
	` //多行字符串
	string_list := []string{"世界", "和平", "真好"}
	fmt.Println("求字符串长度:", len(s1), len(s2))                                                //求字符串长度: 11 13
	fmt.Println("分割字符串:", strings.Split(s1, "l"))                                           //分割字符串: [he  o wor d]
	fmt.Println("判断是否包含：", strings.Contains(s1, "hello"))                                   //判断是否包含： true
	fmt.Println("子串出现位置:", strings.Index(s1, "l"), "子串最后出现位置：", strings.LastIndex(s1, "l")) //子串出现位置: 2 子串最后出现位置： 9
	fmt.Println("join操作:", strings.Join(string_list, "-"))                                  //join操作: 世界-和平-真好
	s3 := "hello世界nihao"
	//因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串                                                                    //因为UTF8编码下一个中文汉字由3~4个字节组成，所以不能简单的按照字节去遍历一个包含中文的字符串
	for _, j := range s3 {
		fmt.Printf("%v =%c  ,", j, j) //104,h 101,e 108,l 108,l 111,o 19990,世 30028,界 110,n 105,i 104,h 97,a 111,o
	}
	res0 := "你好吗"
	res1 := []rune(res0) //强制类型转换 string=>[]rune
	fmt.Println(res1)
	res := []rune{'你', '好', '吗'}
	for i := 0; i < len(res); i++ {
		fmt.Printf("中文输出：%v（%c）\n", res[i], res[i])
	}
	for j := 0; j < len(res1); j++ {
		fmt.Printf("%v:(%c) ", res1[j], res1[j])
	}
	// var m, n int
	// // fmt.Println(m, n)
	// fmt.Scanln(&m, &n)
	// // fmt.Scanln(&m, &n)
	// // fmt.Println(m, n)
	// res := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	res[i] = make([]int, n)
	// }
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Scanln(&res[i][j])
	// 	}
	// }
	// fmt.Println(res)

	// a := strings.IndexByte("abc", 'b')
	// b := strings.Index("abc", "bc")
	// fmt.Println(a, b)
	// c := strings.ToLower("ABCdddd")
	// d := strings.ToUpper("AASDAScccccad")
	// fmt.Println(c, d)
	// e := strings.Contains("Abc", "bd")
	// fmt.Println(e)

	//输入字符串,方式一
	var input1 string
	fmt.Scan(&input1)
	//输入字符串,方式二
	inputReader := bufio.NewReader(os.Stdin) // 使用了自动类型推导，不需要var关键字声明
	input, _ := inputReader.ReadString('a')
	// input2, _ := inputReader.ReadByte()
	// input3, _ := inputReader.ReadBytes(100)
	// input4, _, _ := inputReader.ReadRune()
	fmt.Println(input)

	/*输入数组*/
	// var (
	// 	m int
	// 	n int
	// )
	// fmt.Scan(&m, &n)
	// input := make([][]int, m)
	// for i := 0; i < m; i++ {
	// 	input[i] = make([]int, n)
	// }
	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Scan(&input[i][j])
	// 	}
	// }
	// fmt.Println(input)
	/*输入字节*/
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan() //读取一行内容
	// input1 := input.Bytes()
	// fmt.Println(input1)
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// m, _ := strconv.Atoi(strings.Split(input.Text(), "")[0])

	// fmt.Println(m)
}
func Str_learn() {
	/*获取相邻ab的位置*/
	str := "sabcdba"
	start := 0
	var num byte
	for i := 0; i < len(str); i++ {
		num += str[i]
		if num >= 195 {
			num_len := i - start + 1
			if num == 195 && str[start] < str[i] {
				fmt.Println(start, i, num, num_len, str[start:i+1])
			}

			num -= str[start]
			start++
		}
	}
}
