package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const E = 0.000001
	z := float64(1) //类型转换同下   float64()
	//var z float64 = float64(1)
	k := float64(0)
	for ; ; z = z - (z*z-x)/(2*z) {
		if z-k <= E && z-k >= -E {
			return z
		}
		k = z
	}
}

func main() {
	var s = Sqrt(2)
	fmt.Println(s)
	fmt.Println("运算结果:", math.Pow(s, 2))
	fmt.Println("math.Sqrt(２):", math.Sqrt(2))
}
