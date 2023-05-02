package main

import "fmt"

// use lable to mark the line
// use goto to redirect the execution flow
// to avoid confusing, not recommend to use goto
func goto_lable() {
	fmt.Println("goto_lable:")
	for i := 0; i < 3; i++ {
		if i == 1 {
			goto end
		}
		fmt.Println(i)
	}
end:
}

func main() {
	goto_lable()
}
