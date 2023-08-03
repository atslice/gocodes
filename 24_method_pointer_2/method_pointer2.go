package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 以值为接收者的方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // v 为值，正常调用Vertex的Abs()方法
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs()) // p为指针，也可直接调用Vertex的Abs()方法，实际上其被解释为(*p).Abs()
	fmt.Println(AbsFunc(*p))
	// fmt.Println(AbsFunc(p))  //编译错误，函数AbsFunc()的参数只接收值
}
