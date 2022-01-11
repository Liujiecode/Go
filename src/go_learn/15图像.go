package go_learn

import (
	"fmt"
	"image"
)

func Res15() {

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m)
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
