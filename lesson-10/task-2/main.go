package main

import "fmt"

// TODO 1: Придумай и объяви любую структуру на свой вкус
// Например: Person, Book, Car, Student, Product, Film...
// с 2-3 полями (например: Название, Цена, Год или Имя, Возраст).
type Product struct {
	name  string
	price string
}

// TODO 2: Реализуй для своей структуры интерфейс fmt.Stringer.
// Для этого напиши метод String() string.
// Подсказка: используй fmt.Sprintf("...", ...), чтобы собрать красивую строку с иконками или текстом.
func (p Product) String() string {
	return fmt.Sprintf("Продукт: %s\nЦена: %s", p.name, p.price)

}
func main() {
	// TODO 3:
	// 1. Создай один или два экземпляра своей структуры.
	// 2. Выведи их через обычный fmt.Println(...).
	// 3. Убедись, что fmt.Println сам вызывает твой метод String()!
	apple := Product{name: "1кг. Яблок", price: "100 руб."}

	fmt.Println(apple)
}
