package main

import "fmt"

func getSequence() func() int {
	i := 0
	fmt.Println("inside getSequence(): i = ", i)
	return func() int {
		i += 1
		fmt.Println("inside returned func(): i = ", i)
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 的变量i 为 0 */
	nextNumber := getSequence() // inside getSequence(): i =  0

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber()) //这个执行结果是1
	fmt.Println(nextNumber()) //这个执行结果是2
	fmt.Println(nextNumber()) //这个执行结果是3

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence() //当getSequence()被重新赋值之后，nextNumber1 为一个函数， 函数的变量i值为0; nextNumber函数并不受影响（其变量i的值不受影响）
	fmt.Println(nextNumber1())   //这儿因为是新赋值的，所以是1
	fmt.Println(nextNumber())    // 4, nextNumber不受影响，这个值并没有被销毁，原因就是闭包导致的，尽管外面的函数销毁了，但是内部函数仍然存在，还可以继续走。这个就是闭包
	fmt.Println(nextNumber1())   //新赋值的，继续执行是2
}
