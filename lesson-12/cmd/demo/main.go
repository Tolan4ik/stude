package main

import (
	"fmt"

	// Импортируем наш собственный внутренний пакет!
	"github.com/Tolan4ik/sprint1-demo/internal/user"
)

func main() {
	// TODO 1:
	// Создай пользователя через конструктор user.New:
	u := user.New("Анатолий", 42)
	//
	// Напечатай его имя u.Name и id через u.GetID()
	fmt.Printf("Имя: %s, ID: %d\n", u.Name, u.GetID())
	// Попробуй раскомментировать строчку ниже и посмотри на ошибку компилятора:
	// fmt.Println(u.id) // ❌ Ошибка: u.id undefined (cannot refer to unexported field id)
}
