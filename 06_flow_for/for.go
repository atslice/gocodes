package main

import "fmt"

// check if an int is prime
func is_prime(x int) (bool, int) {
	if x < 2 {
		return false, -1
	} else if x == 2 { // else can not be put at start of a new line. syntax error: unexpected else, expected }
		return true, 0
	} else {
		for tmp := 2; tmp < x; tmp++ { // this works, but stupid
			if x%tmp == 0 {
				return false, tmp
			}
		}
		return true, 0
	}
}

func main() {
	for x := 111; x < 197; x++ {
		// fmt.Printf("is_prime(%d):  " % x)  // Build Error: invalid operation: "is_prime(%d):  " % x (mismatched types untyped string and int) (exit status 1)
		// fmt.Printf("is_prime(%d):  ", % x)  // syntax error: unexpected %, expected expression
		fmt.Printf("is_prime(%d):  ", x) // this works, no need leading  %
		isprime, factor := is_prime(x)
		if !isprime {
			fmt.Println(factor)
		} else {
			fmt.Println("true")
		}
	}
}
