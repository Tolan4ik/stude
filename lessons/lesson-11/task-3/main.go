package main

import (
	"fmt"
	"io"
	"os"
)

// TODO 1: Напиши функцию readFile(name string) ([]byte, error)
func readFile(name string) ([]byte, error) {

	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", name, err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", name, err)
	}
	return data, nil
}

// Шаги внутри функции:
// 1. Открой файл: f, err := os.Open(name)
// 2. Если err != nil, верни: nil, fmt.Errorf("open %s: %w", name, err)
// 3. Поставь отложенное закрытие: defer f.Close()
// 4. Прочитай всё содержимое: data, err := io.ReadAll(f)
// 5. Если err != nil, верни: nil, fmt.Errorf("read %s: %w", name, err)
// 6. Верни: data, nil

func main() {
	// TODO 2:
	// 1. Попробуй прочитать несуществующий файл:

	if data, err := readFile("non_existent.txt"); err != nil {
		fmt.Println(err, data)
	}
	if data, err := readFile("test.txt"); err == nil {
		fmt.Println(string(data))
	}
	// 2. Прочитай существующий файл "test.txt" (я создал его рядом):
	//    data, err := readFile("test.txt")
	//    Если ошибки нет, напечатай его содержимое через string(data).

}
