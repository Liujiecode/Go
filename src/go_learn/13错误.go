package go_learn

import (
	"fmt"
	"math"
	"time"
)

/*实例一：*/
type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

/*实例二*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

/*
error是一个内建接口，自带的
type error interface{
	Error() string
}
*/
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
func Res13() {
	fmt.Println("------实例一：")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println("-----实例二：")
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
