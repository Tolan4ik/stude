package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

func main() {

	res, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
