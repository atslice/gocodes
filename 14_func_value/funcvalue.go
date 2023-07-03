package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // math.Sqrt(5*5 + 12*12)

	fmt.Println(compute(hypot)) // math.Sqrt(3*3 + 4*4), then why not use hypot(3,4) directly? what is its purpose?

	// func math.Pow(x float64, y float64) float64
	// Pow returns x**y, the base-x exponential of y.
	fmt.Println(compute(math.Pow)) // x**y, 3**4 = 3*3*3*3 = 81

	fmt.Println(compute(hypot)) // math.Sqrt(3*3 + 4*4)

	add := func(x, y float64) float64 {
		return x + y
	}

	fmt.Println(compute(add)) // 3 + 4 = 7
}
