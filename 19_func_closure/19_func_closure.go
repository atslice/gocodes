package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	x := 0
	y := 1
	fmt.Println("inside fibonacci(): ", ", x = ", x, ", y = ", y)
	return func() int {
		t := x
		x = y
		y = t + y
		fmt.Println("inside returned func: ", ", x = ", x, ", y = ", y)
		return t
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
