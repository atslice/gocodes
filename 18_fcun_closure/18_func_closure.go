package main

import "fmt"

func appended() func(int) []int {
	a := []int{}
	return func(i int) []int {
		fmt.Println("inside returned func: a = ", a)
		a = append(a, i)
		return a
	}
}

func main() {
	ap := appended()

	fmt.Println(ap(2))
	fmt.Println(ap(5))
	fmt.Println(ap(8))
	result := ap(11)
	fmt.Println(result)
}
