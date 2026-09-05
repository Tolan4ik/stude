package main

import (
	"fmt"
	"strconv"
)

func main() {

	buks := "100"
	// Функция Atoi выдает ДВА значения: число (num) и ошибку (err)
	num, err := strconv.Atoi(buks)

	// Проверяем: если err НЕ РАВНА nil (то есть ошибка произошла)
	if err != nil {
		fmt.Println("Упс! Не удалось перевести в число.")
		fmt.Println("Вот что говорит компьютер:", err)
		return // Останавливаем программу, дальше идти нет смысла
	}

	// Этот код не выполнится, так как программа зайдёт в блок if выше
	fmt.Println("Успех! Число равно:", num)
	buksagain := strconv.Itoa(num)
	fmt.Println(buksagain)
	// Мой первый код в экспериментальной ветке!git
	//письки сиськи

}
