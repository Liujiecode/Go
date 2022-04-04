package leetcode

import (
	"fmt"
	"math"
)

func WangYi01() {
	var (
		a   int
		b   int
		x   int
		y   int
		res int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Println(a, b, x, y)
	if x <= y { //如果y >= x，肯定直接选用群体攻击方式
		max := math.Max(float64(a), float64(b))
		res = (int(max)-1)/y + 1
		fmt.Print(res)
	} else if x <= 2*y { //也就是说群体攻击伤害值加起来比x大，优先使用群体攻击，直到一个怪物死亡，再使用单体攻击攻击另一个怪物
		// min := math.Min(float64(a), float64(b))
		// max := math.Max(float64(a), float64(b))
		// res := (int(min)+1)/y + (int(max)-(int(min)+1)+1)/y
		for a > 0 && b > 0 {
			a -= y
			b -= y
			res += 1
		}
		for a > 0 {
			a -= x
			res += 1
		}
		for b > 0 {
			b -= x
			res += 1
		}
		fmt.Print(res)
	} else { //单体攻击伤害值比群体攻击一共的伤害值还要高，优先使用单体攻击，将每个怪物的生命值打到小于等于y的时候，在使用一个群体攻击直接一起带走
		// res = (a+1)/x + (b+1)/y
		for a > y {
			a -= x
			res++
		}
		for b > y {
			b -= x
			res++
		}
		if a > 0 || b > 0 {
			res++
		}
		fmt.Print(res)
	}
}
