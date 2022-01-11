package go_learn

import (
	"fmt"
	"strings"
)

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
	for _, j := range s3 {
		fmt.Print(j, " ")          //104 101 108 108 111 19990 30028 110 105 104 97 111
		fmt.Printf("%v,%c ", j, j) //104,h 101,e 108,l 108,l 111,o 19990,世 30028,界 110,n 105,i 104,h 97,a 111,o
	}
}
