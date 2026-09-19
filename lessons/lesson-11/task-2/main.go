package main

import (
	"errors"
	"fmt"
)

// TODO 1: Объяви структуру ValidationError с полями:
// - Field string (название поля, например "age")
// - Msg   string (текст ошибки, например "negative")
type ValidationError struct {
	Field string
	Msg   string
}

// TODO 2: Реализуй метод Error() string для указателя (*ValidationError).
// Пусть он возвращает строку вида "Field: Msg" (например, через e.Field + ": " + e.Msg).
func (v *ValidationError) Error() string {
	return v.Field + ": " + v.Msg
}

// TODO 3: Напиши функцию validate(age int) error.
// Если age < 0, верни указатель &ValidationError{Field: "age", Msg: "negative"}.
// Иначе верни nil.
func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "Negative"}
	} else {
		return nil
	}
}

// TODO 4: Напиши функцию createUser(age int) error.
// Она вызывает validate(age).
// Если validate вернула ошибку, оберни её через fmt.Errorf("create user: %w", err) и верни.
// Иначе верни nil.
func createUser(age int) error {
	if err := validate(age); err != nil {
		return fmt.Errorf("create user: %w", err)
	} else {
		return nil
	}
}
func main() {
	// TODO 5:
	// 1. Вызови err := createUser(-5)
	err := createUser(-5)
	// 2. Распакуй ошибку с помощью errors.As:
	//    Объяви переменную-приёмник: var valErr *ValidationError
	//    Проверь через: if errors.As(err, &valErr) { ... }
	//    Внутри напечатай поле с ошибкой: valErr.Field и valErr.Msg
	var valErr *ValidationError
	if errors.As(err, &valErr) {
		// Если внутри матрёшки нашёлся ValidationError:
		fmt.Println("Поле:", valErr.Field)
		fmt.Println("Сообщение:", valErr.Msg)
		// 3. Выведи и саму полную ошибку fmt.Println("Полная ошибка:", err)
		fmt.Println("Полная ошибка целиком:", err)
	}
}
