package main

import (
	"fmt"
	"math/cmplx"
)

// use keyword var to declare a variable, followed by the identifier, and then type of the variable

// variable declaration can be put in package level
var num int // declare a variable named num, which type is int
// fmt.Println(num)  // syntax error: non-declaration statement outside function body

// declare variables in groups
var (
	width  int
	length int
)

// declare multiple variables in a line
var c, python, java bool

// a global variable can be declared but not used
var not_used_var string

// := can only be used inside function body
// size := 100  Build Error: syntax error: non-declaration statement outside function body

func cal(sum int) (x, y int) {
	x = sum / 5
	y = sum - x
	return x, y
}

func main() {
	// variable declaration can be put inside function body
	var i int
	fmt.Println(num)                // a default value of 0 is assigned to the variable of int type
	fmt.Println(i, c, python, java) // a default value of false is assigned to the variable of bool type

	// var not_to_use string  // local variable must be used if it is declared, Buid Error: declared and not used

	var name string = "alice" // assign value in declaration
	fmt.Println(name)

	var age = 20 // assign value in declaration without specifying type
	fmt.Println(age)

	// the variable "age" is assgined value of 20 in declaration, then it is type of int
	// the type of variable can not be changed,
	// age = "Belinda"  // Build Error: cannot use "Belinda" (untyped string constant) as int value in assignment
	fmt.Println(age)

	var a, b int = 5, 6 // assgin values in declaration for multiple variables
	fmt.Println(a, b)

	var s string              // string
	fmt.Println("string:", s) // default value for string is empty string ""

	var f float64             // float64
	fmt.Println("floate:", f) // default value for float is 0

	var pa *int                 // pointer
	fmt.Println("pointer:", pa) // default value for pointer is nil

	var aa []int    // array ?
	fmt.Println(aa) // default value for array is nil ? []

	var ma map[string]int // map ?
	fmt.Println(ma)       // default value for ? is nil map[]

	var ca chan int // chan ?
	fmt.Println(ca) // default value for chan ? is nil

	// var fa func(string) int // func ?  Build Error: declared and not used
	// fmt.Println(fa)         // default value for ? is nil, fmt.Println arg fa is a func value, not called

	var ia error    // interface ?
	fmt.Println(ia) // default value for ? is nil

	fashion := "A new declaration method" // use := to declare a new variable and assign value
	fmt.Println(fashion)

	x, y, z := 3, 4, 5 // declare and assign values for multiple variables by using :=
	fmt.Println(x, y, z)

	xx, yy, zz := 3, "hello", true // declare and assign values for multiple variables of different types by using :=
	fmt.Println(xx, yy, zz)

	width = 800
	length = 600
	fmt.Println("width", width, "length:", length)

	// local variable can be the same name with global variables
	var (
		width  bool
		length int
	)
	fmt.Println("width", width, "length:", length) // the values of local variables, not the global ones

	var lat, long int
	lat, long = 50, 29 // assign values after declaration
	fmt.Println("lat", lat, "long:", long)

	/*
		_ is a writable variable, you can not read it, usually used to discard unneeded values
	*/
	var left int
	_, left = cal(70) // function cal returns two values, but we only needs one, for every variables must be used, use _ to discard unneeded one
	fmt.Println("left:", left)

	// use & to get the memory address of variable
	var weight int = 100
	var height int = 200
	fmt.Println("memory address of variable weight:", &weight, "height:", &height)
	weight = 50
	height = 300
	fmt.Println("memory address of variable weight:", &weight, "height:", &height)
	weight = 100
	height = 200
	fmt.Println("memory address of variable weight:", &weight, "height:", &height)
	height = weight
	fmt.Println("memory address of variable weight:", &weight, "height:", &height)
	// fmt.Println("memory address of 100:", &100)  // Build Error: invalid operation: cannot take address of 100 (untyped int constant)

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		cz     complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", cz, cz)

	// convert type
	var quantity int = 64
	var fq = float64(quantity)
	var uq uint = uint(fq)
	fmt.Println("quantity:", quantity, "float64:", fq, "unint:", uq)

	var quality int = 258
	var fql = float64(quality)
	var uql uint8 = uint8(fql) // the range for uint8 is from 0 to 255, thus the converted value for 256 is 0, for 257 is 1, for 258 is 2, etc
	// var u257 uint8 = uint8(257)  // Build Error: cannot convert 257 (untyped int constant) to type uint8
	fmt.Println("quality:", quality, "float64:", fql, "unint8:", uql)

}
