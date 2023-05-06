package main

import "fmt"

// defer statement will not be executed until outer function is executed
func defer_fun() {
	defer fmt.Println("defer statement of defer_fun") // defer statement will not be executed until outer function is executed
	fmt.Println("defer_fun:")
}

// defer execution order: first in last out
func defer_order() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer statement: ", i) // defer statement will not be executed until outer function is executed
	}
	fmt.Println("defer_order:")
}

func main() {

	defer_fun()
	defer_order()

}
