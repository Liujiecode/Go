package go_learn

import (
	"fmt"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

//方法:值接收者方法,一般只接收值不更改，因为更改也没有用，不会影响真正的接收者
func (s Student) Get() int {
	return s.ID
}
func (s Student) Get2(id int) int {
	s.ID = id
	return s.ID
}

//方法：指针接收者方法，可用来修改值，更改会影响真正的接收者
func (s *Student) Set(id int) {
	s.ID = id
}

//函数
func Notify2(id int) {
	fmt.Println("学生id为：", id)
}
func Res07() {
	s1 := Student{}
	fmt.Println("s1为：", s1)
	fmt.Println("方法——指针接收者：")
	fmt.Println("使用值接收者方法改变后id：", s1.Get2(16), "实际的id没变：", s1.ID)
	fmt.Println("使用指针接收者方法前：", s1.Get())
	s1.Set(20)
	fmt.Println("使用指针接收者方法后：", s1.Get())
	s2 := &Student{}
	fmt.Println("使用指针接收者方法前：", s2.Get())
	s2.Set(30)
	fmt.Println("使用指针接收者方法后：", s2.Get())

	fmt.Print("函数—— ：")
	Notify2(17)

}
