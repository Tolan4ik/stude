package main

import "fmt"

// TODO 1: Объяви интерфейс Greeter с одним методом Greet() без параметров и без возвращаемых значений.
type Greeter interface {
	Greet() string
}

// TODO 2: Создай структуру EN (английский) и добавь ей метод Greet(), который выводит "Hello!" в консоль.
type EN struct {
	e string
}

func (e EN) Greet() string {
	return "Hello"
}

// TODO 3: Создай структуру RU (русский) и добавь ей метод Greet(), который выводит "Привет!" в консоль.
type RU struct {
	r string
}

func (r RU) Greet() string {
	return "Привет"
}

// TODO 4: Напиши функцию Say, которая принимает параметр типа Greeter и вызывает у него метод Greet().

func say(g Greeter) string {
	return g.Greet()
}

func main() {
	// TODO 5:
	// 1. Создай экземпляры структур EN и RU.
	en := EN{}
	ru := RU{}
	// 2. Вызови функцию Say для каждого из них.
	se := say(en)
	sr := say(ru)
	// 3. (По желанию) Создай срез []Greeter, положи туда обе структуры и вызови метод Greet() в цикле for range.

	fmt.Printf("Английский: %s \nРусский: %s\n", se, sr)
}
