package main

import "fmt"

func main() {
	testScore := 10

	// Вызываем функцию grade, передавая в неё значение переменной testScore (10).
	// Результат работы функции (строку "D") сохраняем в переменную finalGrade.
	finalGrade := grade(testScore)

	// Выводим полученный результат на экран: "Оценка за баллы: D"
	fmt.Println("Оценка за баллы:", finalGrade)
}

// Объявляем функцию с именем grade.
// Она принимает один аргумент (testScore) типа int и ОБЯЗАНА вернуть значение типа string.
func grade(testScore int) string {
	// Используем switch без выражения в качестве замены цепочки if-else
	switch {
	// Проверяем сверху вниз. Первое же условие, давшее true, вызовет return
	case testScore >= 90:
		// Оператор return мгновенно завершает работу функции и возвращает значение.
		// Писать break здесь не нужно.
		return "A"
	case testScore >= 80:
		return "B"
	case testScore >= 70:
		return "C"
	default:
		// Если ни одно из условий выше не подошло (как в нашем случае, ведь 10 < 70),
		// срабатывает ветка по умолчанию (default).
		return "D"
	}
}
