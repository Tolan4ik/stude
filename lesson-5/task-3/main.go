package main

import "fmt"

func apply(n int, fn func(int) int) int {
	return fn(n)
}

func main() {
	res := apply(5, func(x int) int { return x * x })
	fmt.Println(res)
}
