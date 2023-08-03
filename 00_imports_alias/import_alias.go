package main

// 通过使用包的别名来解决包名之间的名称冲突
import fm "fmt" // alias3

func main() {
	fm.Println("hello, world")
}
