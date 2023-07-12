package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Closures can also be recursive, but this requires the closure to be declared with a typed var explicitly before it’s defined.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	// Since fib was previously declared in main, Go knows which function to call with fib here.
	fmt.Println(fib(7))
	for i := -5; i < 10; i++ {
		fmt.Println(fib(i))
	}
}
