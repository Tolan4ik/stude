package main

import "fmt"

// Создадим простую структуру для проверки (можно использовать свою)
type Product struct {
	name string
}

// TODO 1: Напиши функцию describe(x any) string
// Внутри используй switch v := x.(type) { ... }
// Обработай кейсы:
// - int: верни строку через fmt.Sprintf("int: %d", v)
// - string: верни fmt.Sprintf("string: %s", v)
// - bool: верни fmt.Sprintf("bool: %t", v)
// - Product: верни fmt.Sprintf("product: %s", v.name)
// - default: верни "unknown"
func describe(x any) string {
	switch v := x.(type) {
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return fmt.Sprintf("string: %s", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	case Product:
		return fmt.Sprintf("product: %s", v.name)
	default:
		return "unknown"
	}

}

func main() {
	// TODO 2: Проверь функцию describe на разных типах:
	fmt.Println(describe(42))
	fmt.Println(describe("hi"))
	fmt.Println(describe(true))
	fmt.Println(describe(Product{name: "Кофе"}))
	fmt.Println(describe(3.14)) // float64 нет в switch, должно вернуть "unknown"

	fmt.Println("--- Часть 1: describe ---")

	fmt.Println("\n--- Часть 2: Type Assertion (v, ok) ---")
	// TODO 3:
	// 1. Создай переменную var a any = 100
	// 2. Достань int через num, ok := a.(int) и выведи num и ok в консоль.
	//
	// 3. Создай переменную var b any = "не число"
	// 4. Попробуй достать int через n, ok := b.(int) и выведи n и ok.
	//    Посмотри, что будет в ok, когда тип не совпал!
	var a any = 100
	if num, ok := a.(int); ok {
		fmt.Println(num, ok)
	}
	var b any = "Не число"
	if num, ok := b.(int); ok {
		fmt.Println(num, ok)
	} else {
		fmt.Println("Это не число")
	}
}
