package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	// diff = 0.001
	// assume = 1
	// 1.abs(x - assume ** 2) < diffかを確認
	// 2.異なる場合は、Adjestを行う。
	// 3.あっている場合は、assumeを返す。
	const (
		threshold = 0.001
	)
	assume := 1.0

	for {
		diff := math.Abs(x - (assume * assume))

		if threshold >= diff {
			return assume
		}

		fmt.Println("%V, %V", diff, assume)

		assume = adjust(assume, x)
	}
}

func adjust(z float64, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func main() {
	fmt.Println(sqrt(2))
}
