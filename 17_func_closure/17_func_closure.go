package main

import "fmt"

// 闭包使用方法
func add(x1, x2 int) func() (int, int) {
	i := 0
	fmt.Println("inside add(): i = ", i)
	return func() (int, int) {
		i++
		fmt.Println("inside returned func(): i = ", i, ", x1 = ", x1, ", x2 = ", x2)
		return i, x1 + x2 // variables i, x1, x2 are outside the scope of the function closure
	}
}

func main() {
	add_func := add(7, 8)   // the variable inside function add() is assigned value of 0, the anonymous func() (int, int) is not executed
	fmt.Println(add_func()) // the function add_func()
	fmt.Println(add_func())
	fmt.Println(add_func())
}
