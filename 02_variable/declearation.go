package main

import "fmt"

// use keyword var to declare a variable, followed by the identifier, and then type of the variable

// variable declaration can be put in package level
var num int // declare a variable named num, which type is int
// fmt.Println(num)  // syntax error: non-declaration statement outside function body

// declare multiple variables in a line
var c, python, java bool

func main() {
	// variable declaration can be put inside function body
	var i int
	fmt.Println(num)                // a default value of 0 is assigned to the variable of int type
	fmt.Println(i, c, python, java) // a default value of false is assigned to the variable of bool type

	var name string = "alice" // assign value in declaration
	fmt.Println(name)

	var age = 20 // assign value in declaration without specifying type
	fmt.Println(age)

	// the variable "age" is assgined value of 20 in declaration, then it is type of int
	// the type of variable can not be changed,
	// age = "Belinda"  // Build Error: cannot use "Belinda" (untyped string constant) as int value in assignment
	fmt.Println(age)
}
