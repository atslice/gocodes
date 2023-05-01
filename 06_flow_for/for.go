package main

import "fmt"

// typical usage of for:
// for init; condition; post {...}
func for_init() {
	fmt.Println("for_init:")
	for i := 0; i < 3; i++ {
		// fmt.Printf(i)  // Build Error: cannot use i (variable of type int) as string value in argument to fmt.Printf (exit status 1)
		fmt.Println(i)
	}
	// the declared variable in for initialization, is only available in for body
	// fmt.Println(i)  // Build Error: undefined: i (exit status 1)
}

// init and post are not must-required
// for condition {...} is similar with while in other language
func for_condition() {
	fmt.Println("for_condition:")
	i := 1
	// for ;i < 3; {...} is also OK
	// for i < 3; {...}  // syntax error: unexpected {, expected for loop condition
	// for ;i < 3 {...} // unexpected {, expected semicolon or newline
	for i < 3 {
		i += i
	}
	fmt.Println(i)
}

// for body will always be executed without condition, which is similar with for true
// use key word break to end the for loop
func for_break() {
	fmt.Println("for_break:")
	i := 1
	for {
		i += i
		if i > 3 {
			break
		}
	}
	fmt.Println(i)
}

// use key word continue to skip the rest logic and go to next loop
func for_continue() {
	fmt.Println("for_continue:")
	sum := 0
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // skip even
		}
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}

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

func print_prime() {

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

func main() {
	for_init()
	for_condition()
	for_break()
	for_continue()

}
