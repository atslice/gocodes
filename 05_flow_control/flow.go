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

func abs(x float64) float64 {
	// unlike Python one-sentence code
	// return x if x > 0 else -x  // Build Error: syntax error: unexpected if at end of statement
	// if x > 0 {return x}; else {return -x} // Build Error: syntax error: unexpected else, expected }
	// if x > 0 {return x} else {return -x}  // This works

	if x > 0 {
		return x
	} else //   Build Error: syntax error: unexpected ) at end of statement (exit status 1)
	{
		return -x
	}

}

func main() {
	fmt.Println(weekday(2023, 04, 27))
	fmt.Println("min(89, 78):", min(89, 78))
	fmt.Println("min(3.98, 3.9799):", min(3.98, 3.9799))
	fmt.Println("abs(-5.8):", abs(-5.8))
	fmt.Println("abs(5.8):", abs(5.8))
}
