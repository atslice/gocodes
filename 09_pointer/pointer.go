package main

import "fmt"

// the value of a pointer is the memory address
func pointer_declare() {
	fmt.Println("pointer_declare():")
	var pt *int // declare a pointer, the zero value is nil
	i := 15
	pt = &i // the & operator gets the memory address of the variable
	fmt.Println(pt)
	fmt.Println(*pt)
}

// "runtime error: invalid memory address or nil pointer dereference"
func pointer_declare_error() {
	var pt *int // declare a pointer, the zero value is nil
	*pt = 15    // "runtime error: invalid memory address or nil pointer dereference"
	fmt.Println(pt)
	fmt.Println(*pt)
}

// the type of the variable that a pointer points to much match the pointer declaration;
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

func pointer() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

func main() {
	pointer_declare()
	pointer()
	pointer_assignment_error()
}
