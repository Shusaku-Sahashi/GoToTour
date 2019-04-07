package main

import (
	"fmt"
	"math"
)

type errorNegativeSqrt float64

func (e errorNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	// diff = 0.001
	// assume = 1
	// 1.abs(x - assume ** 2) < diffかを確認
	// 2.異なる場合は、Adjestを行う。
	// 3.あっている場合は、assumeを返す。
	if x < 0 {
		return 0, errorNegativeSqrt(x)
	}

	const (
		threshold = 0.001
	)
	assume := 1.0

	for {
		diff := math.Abs(x - (assume * assume))

		if threshold >= diff {
			return assume, nil
		}

		assume = adjust(assume, x)
	}
}

func adjust(z float64, x float64) float64 {
	return z - (z*z-x)/(2*z)
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
