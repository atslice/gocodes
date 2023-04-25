package main

import "fmt"

func add(x, y int) int {
	return x + y // add: +
}

func minus(x, y float64) float64 {
	return x - y // minus: -
}

func greater(x, y float64) bool {
	return x > y
}

func smaller(x, y float64) bool {
	return x < y
}

func equal(x, y float64) bool {
	return x == y
}

func main() {
	fmt.Println(add(5, 20))
	fmt.Println(minus(598.6, 45.26))
	fmt.Println("9 is greater than 6:", greater(9, 6))
	fmt.Println("9 is greater than 9.0:", greater(9, 9.0))
	fmt.Println("9 is smaller than 9.0:", smaller(9, 9.0))
	fmt.Println("9 is equal to 9.0:", equal(9, 9.0))
}
