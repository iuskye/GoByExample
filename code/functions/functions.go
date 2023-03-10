package main

import "fmt"

// 这里是一个函数，接受两个 int 并且以 int 返回它们的和。
func plus(a int, b int) int {
	return a + b
}

// Go 需要明确的 return，也就是说，它不会自动 return 最后一个表达式的值。
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// 通过函数名(参数列表)来调用函数。
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
