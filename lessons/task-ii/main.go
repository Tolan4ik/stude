package main

import "fmt"

func main() {

	numbers := []int{10, 15, 22, 33, 40, 55}

	for x, y := range numbers {
		if y%2 == 0 {
			fmt.Printf("\nЧётное число: %d \nи его индекс: %d", y, x)
		}
	}

}
