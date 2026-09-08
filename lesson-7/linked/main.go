package main

import "fmt"

type Step struct {
	Name string
	Next *Step
}

func main() {

	step3 := &Step{
		Name: "3. Отправить чек клиенту",
	}
	step2 := &Step{
		Name: "2. Списать 5000 тенге",
		Next: step3,
	}
	step1 := &Step{
		Name: "1. Проверить баланс",
		Next: step2,
	}

	for current := step1; current != nil; current = current.Next {
		fmt.Println(current.Name)
	}
}
