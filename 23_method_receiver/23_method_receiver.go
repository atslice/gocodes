package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// the parameter p before the method name is the receiver of the method Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// define another method for type Point, calculate the distance from zero Point{0, 0}
func (p Point) DistanceZero() float64 {
	return p.Distance(Point{0, 0}) // use the method we just define
}

// operate the copied value
func (v Point) ValueScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// pp is a pointer which points to type Point
func (pp *Point) Scale(f float64) {
	pp.X = pp.X * f
	pp.Y = pp.Y * f
}

func main() {
	p := Point{3, 4}
	fmt.Println(p.DistanceZero()) //
	fmt.Println(p)

	p.ValueScale(10)              // does not modify the value of p
	fmt.Println(p.DistanceZero()) //
	fmt.Println(p)

	p.Scale(10)                   // deos modify the value of p
	fmt.Println(p.DistanceZero()) //
	fmt.Println(p)

	p.ValueScale(10)              // does not modify the value of p
	fmt.Println(p.DistanceZero()) //
	fmt.Println(p)

}
