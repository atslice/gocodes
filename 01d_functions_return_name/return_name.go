package main

import "fmt"

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
	fmt.Println(split(17))
	fmt.Println(split2(17))
}
