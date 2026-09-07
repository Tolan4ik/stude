package main

import "fmt"

func apply(n int, fn func(int) int) int {
	return fn(n)
}

func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	res := apply(5, func(x int) int { return x * x })
	fmt.Println(res)

	re := calc(10, 5, func(x, y int) int { return x + y })
	umn := calc(10, 5, func(x, y int) int { return x * y })
	fmt.Println(re)
	fmt.Println(umn)
}
