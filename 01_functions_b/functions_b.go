package main

import "fmt"

/*
consequent arguments for functions are the same type, only the last argument needs declearation
	funct add(x, y int) int
which is the same with:
	func add(x int, y int) int
*/

func add(x, y int) int {
	return x + y
}

func sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sum(42, 13))
}
