package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method of Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method of Vertex, pointer as receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function, the parameter v is a pointer
// if called, the passed parameter v must be a pointer but not a value
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	v1.Scale(10)          // though v1 is value, and the method Scale appect a pointer as receiver parameter, it is OK
	fmt.Println(v1.Abs()) // 50
	(&v1).Scale(10)       // pointer is also OK
	fmt.Println(v1.Abs()) // 500

	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	// Scale(v, 10)  // cannot use v (variable of type Vertex) as *Vertex value in argument to Scale
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
