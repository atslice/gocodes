package main

import (
	"fmt"
	"unsafe"
)

// use key word const to define constant
const Pi = 3.1415

// const identifier [type] = value
const Name string = "Alice"

// or ignore type
// const identifier = value
const Height = 100

// define constants in a line
const ip, port = "127.0.0.1", 5638

// define constants in group, for example, enum
const (
	Anoy        = 1
	High        = 0
	Transparent = 2
	Unknown     = 3
)

// import "unsafe"  // Build Error: imports must appear before other declarations

// constant can be declared by built-in function, such as len(), cap(), unsafe.Sizeof()

const (
	sentence = "This is a sentence."
	length   = len(sentence)           // 19
	size     = unsafe.Sizeof(sentence) // 16
)

/*

It is said that constant can not be declared by user's function, otherwise build-error.

func get_int() int {
	return 100
}

const haha = get_int()  // get_int() (value of type int) is not constant

func get_Height() {
	return
}

const hh = get_Height()  // get_Height() (no value) used as value
*/

func get_int() int {
	return 100
}

var bint = get_int()

const (
	a = 2
	b // in group declaration of constants, declared constant without value, the value of the last line will be used as the default value
	c
	d
)

const aa = 3

// can not miss expression in non-group declaration of constant
// const bb  // Build Error:  missing init expr for bb
// const cc  // Build Error:  missing init expr for cc

func main() {
	fmt.Println("Pi: ", Pi)
	fmt.Println("Name: ", Name)
	fmt.Println("Hight: ", Height)
	fmt.Println("ip: ", ip)
	fmt.Println("port: ", port)
	fmt.Println("sentence:", sentence, "length:", length, "size:", size)
	fmt.Println("bint:", bint)
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d)
	fmt.Println("aa:", aa)
}
