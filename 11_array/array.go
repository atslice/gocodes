package main

import "fmt"

// array stores a group of variables of the same type
// var identifier [n]type
// where n is int, the number of the members, it is fixed
func array_declare() {
	fmt.Println("array_declare()")
	var a [2]string // declare an array of string identified as a, which has two memebers
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a) // [Hello World]

	var integer []int    // an empty array
	fmt.Println(integer) // []

	// var ff [3]float64 = [3.14, 5.25, 6.26]  // Build Error: syntax error: unexpected comma; expected ] (exit status 1)
	var ff [5]float64 = [5]float64{3.14, 5.25, 6.26} // zero-value will be given to non-initialized items
	fmt.Println(ff)                                  // [3.14 5.25 6.26 0 0]

	// var fff [5]float64 = [3]float64{3.14, 5.25, 6.26}  //  Build Error: cannot use [3]float64{…} (value of type [3]float64) as [5]float64 value in variable declaration (exit status 1)
	var fff [5]float64 = [5]float64{}
	fmt.Println(fff) // [0 0 0 0 0]

	var f5 [5]float64 = [5]float64{0: 8.29, 2: 5.78} // initialize by index
	fmt.Println(f5)                                  // [8.29 0 5.78 0 0]
}

// array can be initialized while declaration
func array_init() {
	fmt.Println("array_init()")
	primes := [6]int{2, 3, 5, 7, 11, 13} // initialize while declaration
	fmt.Println(primes)                  // [2 3 5 7 11 13]
}

// let compiler count the members when initialized
func array_init2() {
	fmt.Println("array_init2()")
	a := [...]string{"hello", "world"}
	fmt.Println(a)
}

func main() {

	array_declare()
	array_init()
	array_init2()

}
