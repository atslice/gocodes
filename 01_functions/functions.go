package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

/*
consequent arguments for functions are the same type, only the last argument needs declearation

	funct add(x, y int) int

which is the same with:

	func add(x int, y int) int
*/
func sum(x, y int) int {
	return x + y
}

/*
function can return multiple values
*/
func swap(x, y string) (string, string) {
	return y, x
}

/*
below is an example of direct return, i.e without var/values after return statement
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*
direct return is not recomanded in complicated functions
*/
func split2(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(sum(42, 13))

	a, b := swap("hello", "world") // return multiple value
	fmt.Println(a, b)

	fmt.Println(split(17)) // direct return
	fmt.Println(split2(17))
}
