// package designmodel
package main

import (
	"fmt"
	"sync"
)

type OutPutData struct {
	data interface{}
}

var data *OutPutData
var once sync.Once

//该类型可以实现单例模式，虽然是多次赋值，但是只执行一次（一个对象多次实例化，但是只有一个，共享对象地址）
func GetInstance(i int) *OutPutData {
	once.Do(func() {
		data = &OutPutData{i}
	})
	return data
}
func main() {
	i1 := GetInstance(1)
	i2 := GetInstance(2)
	if i1 == i2 {
		fmt.Println("单例模式")
	} else {
		fmt.Println("不是单例模式")
	}
}
