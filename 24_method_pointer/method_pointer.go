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

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{3, 4}
	fmt.Println(v1.Abs())
	v1.Scale(10)
	fmt.Println(v1.Abs())

	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
