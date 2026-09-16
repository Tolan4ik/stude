package main

import "fmt"

func main() {

	m := map[string]int{
		"alice":   30,
		"bob":     25,
		"Tola":    25,
		"Olga":    64,
		"Tana":    43,
		"Natasha": 24,
	}
	delete(m, "alice")

	_, ok := m["alice"]
	if !ok {
		fmt.Println("Нет ключа")

	}

	fmt.Println("Первый круг")
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println("Второй круг")
	for key, value := range m {
		fmt.Println(key, value)
	}
}
