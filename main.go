package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	a := 2
	b := 10
	c := 4
	d := 6
	f := 9.0
	e := 5.0
	fmt.Println(a + b)
	fmt.Println(b - a)
	fmt.Println(b * a)
	fmt.Println(b / a)
	fmt.Println(d / c)
	fmt.Println(d % c)
	fmt.Println(f / e)
	j := "Привет"
	fmt.Println(j)
	fmt.Println(len(j))
	l := "ржака😂"
	fmt.Println(l)
	fmt.Println(l[10])
	fmt.Println(len(l))

	s := "Привет"

	for i, r := range s {
		fmt.Printf("%d  %c\n", i, r)
	}

	runes := []rune(s)
	fmt.Println(string(runes[0])) // "П"
	fmt.Println(len(runes))       // 6 символов
	fmt.Printf("%c\n", runes[0])
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Print(`hi
	bye
	`)
	runes[0] = 'Ы'
	s = string(runes)
	fmt.Println(s)
	hi := "hellllo"
	he := []byte(hi)
	he[0] = 'H'
	hi = string(he)
	fmt.Println(hi)
	var ok bool
	fmt.Println(ok)
	var n int
	var o string
	fmt.Println(n)
	fmt.Println(o)
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

	fmt.Println("калякамаляка тут в ветка а")

	fmt.Println("маляка каляка вот такое тут в ветке Б елы палы маталы")
	fmt.Println("что тут за рамсы что за конфликты не дерёмся")
}
