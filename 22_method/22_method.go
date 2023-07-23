package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
// the parameter p before the method name is the receiver of the method Distance
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// define another method for type Point, calculate the distance from zero Point{0, 0}
func (p Point) DistanceZero() float64 {
	return p.Distance(Point{0, 0}) // use the method we just define
}

// 在Go语言里，我们为一些简单的数值、字符串、slice、map来定义一些附加行为很方便。
// 我们可以给同一个包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者interface。
type MyFloat float64 // MyFloat为命名类型，底层类型为float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))   // "5", function call
	fmt.Println(p.Distance(q))    // "5", method call
	fmt.Println(p.DistanceZero()) // 2.23606797749979
	fmt.Println(q.DistanceZero()) // 7.211102550927979

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	g := MyFloat(-5.0)
	fmt.Println(g.Abs())

}
