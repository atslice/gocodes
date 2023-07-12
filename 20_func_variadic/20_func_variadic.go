package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.
// Within the function, the type of nums is equivalent to []int.
// We can call len(nums), iterate over it with range, etc.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func con(words ...string) {
	sum := ""
	for _, word := range words {
		if sum != "" {
			sum = sum + "-" + word
		} else {
			sum = word
		}
	}
	fmt.Println(sum)
}

func main() {

	sum(1, 2) // Variadic functions can be called in the usual way with individual arguments.
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.

	con("Hello", "World")
	con("Nobody", "wana", "to", "be", "defeated")
	words := []string{"Never", "grow", "old"}
	con(words...)
}
