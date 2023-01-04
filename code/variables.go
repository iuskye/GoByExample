package main

import "fmt"

func main() {
	// var 声明 1 个或者多个变量
	var a = "initial"
	fmt.Println(a)

	// 你可以一次性声明多个变量。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go 会自动推断已经有初始值的变量的类型。
	var d = true
	fmt.Println(d)

	// 声明后却没有给出对应的初始值时，变量将会初始化为零值。例如，int 的零值是 0。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写，例如：var f string = "apple" 可以简写为下边这样。
	f := "apple"
	fmt.Println(f)
}
