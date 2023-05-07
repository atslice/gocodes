package main

import "fmt"

// the value of a pointer is the memory address
func pointer_declare() {
	fmt.Println("pointer_declare():")
	var pt *int // declare a pointer, the zero value is nil
	i := 15
	pt = &i          // the & operator gets the memory address of the variable
	fmt.Println(pt)  // memory address
	fmt.Println(*pt) // read the value of variable i by pointer
}

// "runtime error: invalid memory address or nil pointer dereference"
func pointer_declare_error() {
	var pt *int // declare a pointer, the zero value is nil
	*pt = 15    // "runtime error: invalid memory address or nil pointer dereference"
	fmt.Println(pt)
	fmt.Println(*pt)
}

// operate variable by pointer
func pointer() {
	fmt.Println("pointer():")
	j := 56
	p := &j         // point to j
	*p = *p / 29    // operate j by pointer
	fmt.Println(*p) // check the value
	fmt.Println(j)  // check the value
}

// the type of the variable that a pointer points to must match the pointer declaration;
func pointer_assignment_error() {
	var pti *int // declare a pointer, the zero value is nil
	var pts *string
	i := 15
	pti = &i
	s := "pointer_assignment_error"
	// pti = &s // the type of the variable that a pointer points to much match the pointer declaration;  Build Error: cannot use &s (value of type *string) as *int value in assignment (exit status 1)
	pts = &s
	fmt.Println(pti)
	fmt.Println(pts)
	fmt.Println(*pts)
}

// go does not support operation of pointer
// question: what is the purpose of introducing pointer? What is the use case?

func main() {
	pointer_declare()
	pointer()
	pointer_assignment_error()
}
