package go_learn

import (
	"fmt"
	"io"
	"strings"
)

//实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。

func Res14() {

	r := strings.NewReader("Hello, Reader!") //字符串长14个字节
	b := make([]byte, 14)                    //创建8字节长度的数组，会循环3次，直到以io.EOF开头循环停止；创建14字节长度的数组，会循环2次；
	fmt.Println("创建初始数组：", b)
	for {
		n, err := r.Read(b) //Read用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
		fmt.Printf("n=%v err=%v b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

}
