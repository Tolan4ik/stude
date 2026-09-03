package main

import (
	"errors"
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

	text := "Я учу Go!"
	numy := 0
	for _, r := range text {

		if r == 'у' {
			numy++
		}

	}
	fmt.Println(numy)

	password := "Super$Sec%ret"

	for _, run := range password {
		if run == '$' || run == '%' {
			fmt.Println("Найден запрещенный символ!")
			break
		}

	}
	fmt.Println("изи проверочка")

	text2 := "Привет! Как твои дела? #программирование"

	for in, rr := range text2 {
		if rr == '#' {
			fmt.Printf("Обнаружен хэштег на позиции %d!", in)
			break
		}
	}
	fmt.Println("\nСканирование текста завершено")
	fmt.Println("--- Старт программы ---")

	// 1. Вызываем функцию и создаем переменную err, куда прилетит результат.
	// Знак := сам создаст переменную err типа error.
	er := checkTemperature(25)

	// 2. Проверяем, что к нам прилетело
	if er != nil {
		// Если прилетела ошибка (не nil), пишем об этом
		fmt.Println("Упс! Функция вернула ошибку:", err)
	} else {
		// Если прилетел nil, значит всё прошло успешно
		fmt.Println("Ура! Функция вернула nil. Температура в норме!")
	}

	fmt.Println("--- Конец программы ---")
}

// Функция проверяет температуру.
// Она обещает вернуть результат типа error (ошибку).
func checkTemperature(temp int) error {
	if temp > 100 {
		// Если температура слишком высокая, создаем и возвращаем ошибку.
		// Она полетит в main() в переменную err.
		return errors.New("перегрев! Температура выше 100 градусов")
	}

	// Если всё хорошо (температура 100 или меньше),
	// мы возвращаем встроенное слово nil (означает "ошибки нет").
	// Этот nil точно так же полетит в main() в переменную err.
	return nil
}
