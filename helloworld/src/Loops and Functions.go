package main

import (
	"fmt"
	"math"
)

func Squrt(x float64) float64 {
	// diff = 0.001
	// assume = 1
	// 1.abs(x - assume ** 2) < diffかを確認
	// 2.異なる場合は、Adjestを行う。
	// 3.あっている場合は、assumeを返す。
	const (
		threshold = 0.001
	)
	var assume float64 = 1.0

	for {
		diff := math.Abs(x - (assume * assume))

		if threshold >= diff {
			return assume
		}

		fmt.Println("%V, %V", diff, assume)

		assume = Adjust(assume, x)
	}
}

func Adjust(z float64, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func main() {
	fmt.Println(Squrt(2))
}
