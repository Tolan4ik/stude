package main

import "fmt"

func GetOrCompute(cache map[string]int, key string, compute func() int) int {
	// 1. Проверь наличие key в cache через v, ok := ...
	//    Если ok == true, сразу верни v.

	v, ok := cache[key]
	if ok {
		return v

	}

	res := compute()
	cache[key] = res
	return res

}

// 2. Если ключа нет, вызови функцию compute() и запиши результат в переменную.

// 3. Сохрани полученный результат в cache под ключом key.

// 4. Верни этот результат.

func main() {

	cache := make(map[string]int)

	expensiveCalc := func() int {
		fmt.Println("--> Выполняем тяжелое вычисление...")
		return 42
	}
	fmt.Println("Вызов 1:", GetOrCompute(cache, "answer", expensiveCalc))
	fmt.Println("Вызов 2:", GetOrCompute(cache, "answer", expensiveCalc))
}
