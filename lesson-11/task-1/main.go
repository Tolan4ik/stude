package main

import (
	"errors"
	"fmt"
)

// TODO 1: Объяви глобальную переменную ErrEmpty с ошибкой "empty".
// Используй функцию errors.New("empty").

var ErrEmpty = errors.New("empty")

// TODO 2: Напиши функцию read(name string) error
// Логика:
// - Если name == "", верни ErrEmpty.
// - Иначе верни nil (ошибки нет).
func read(name string) error {
	if name == "" {
		return ErrEmpty
	} else {
		return nil
	}
}

func main() {
	// TODO 3:
	// 1. Вызови read("") и сохрани результат в переменную err.
	// 2. С помощью if errors.Is(err, ErrEmpty) проверь ошибку
	//    и напечатай, например: "Ошибка: передано пустое имя!".
	//
	// 3. Вызови read("document.txt") с непустым именем
	//    и убедись, что err == nil (всё прошло успешно).

	if err := read(""); errors.Is(err, ErrEmpty) {
		fmt.Println("Ошибка: передано пустое имя!")
	}
	if err := read("document.txt"); err == nil {
		fmt.Println("Успех")
	}

}
