package main

import (
	"fmt"
	"math/rand"
)

func contains(s []int, target int) bool {
	for _, val := range s {
		if val == target {
			return true // нашли!
		}
	}
	return false
}

// uniqN возвращает срез из n уникальных случайных чисел.
func uniqN(n int) []int {
	res := make([]int, 0, n)
	for len(res) < n {
		num := rand.Intn(100)
		if !contains(res, num) {
			res = append(res, num)
		}

	}
	return res
}

func main() {
	res := uniqN(5)
	fmt.Println("Результат:", res)
}
