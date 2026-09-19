package main

import "fmt"

func main() {
	role := "admin"

	// Здесь switch проверяет конкретное значение переменной role.
	// Go по очереди сравнивает строку "admin" с каждым вариантом в case.
	switch role {
	// Сравниваем: "admin" == "user"? Ответ: false. Идем дальше.
	case "user":
		fmt.Println("Нет доступа")

	// Сравниваем: "admin" == "moderator"? Ответ: false. Идем дальше.
	case "moderator":
		fmt.Println("Ограниченый доступ")

	// Сравниваем: "admin" == "admin"? Ответ: true!
	// Этот блок выполняется. Благодаря встроенному break, после выполнения
	// этого case управление сразу выходит из switch.
	case "admin":
		fmt.Println("\nПолный доступ") // \n сделает пустую строку перед текстом

	// Сработает только если значение role не совпало ни с одним case
	// (например, если бы role была равна "guest")
	default:
		fmt.Println("Неизвестная роль")
	}
}
