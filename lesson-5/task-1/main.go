package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

func area(a, b int) (int, error) {
	if a <= 0 || b <= 0 {
		return 0, fmt.Errorf("стороны должны быть больше нуля")
	}
	return a * b, nil
}

func withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return 0, fmt.Errorf("сумма снятия должна быть больше нуля")
	} else {
		if amount > balance {
			return 0, fmt.Errorf("Недостаточно средств")
		}
	}
	return balance - amount, nil
}

func main() {

	res, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	re, er := area(2, 4)
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println(re)
	}

	resu, ero := withdraw(500, 100)
	if ero != nil {
		fmt.Println(ero)
	} else {
		fmt.Println(resu)
	}

}
