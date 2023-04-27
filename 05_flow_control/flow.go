package main

import "fmt"

// give what the weekday is for specified date
func weekday(year, month, day int) int {
	return 0
}

// if statement example
// () is not needed for the condition statement
func min(x, y float64) float64 {
	if x > y {
		return y
	} else {
		return x
	}
}

func main() {
	fmt.Println(weekday(2023, 04, 27))
	fmt.Println("min(89, 78):", min(89, 78))
	fmt.Println("min(3.98, 3.9799):", min(3.98, 3.9799))
}
