package main

import "fmt"

func max(nums ...int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("список пуст")
	}
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n

		}
	}

	return m, nil
}

func main() {
	maxim, err := max(1, 5, 99, 3, 42)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maxim)
	}
}
