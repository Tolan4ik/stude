package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-manager/internal/storage"
)

func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readInt(reader *bufio.Reader, prompt string) (int, error) {
	str := readLine(reader, prompt)
	return strconv.Atoi(str)
}

func clearScreen() {
	// ANSI-код очистки экрана
	fmt.Print("\033[H\033[2J")
}

func waitEnter(reader *bufio.Reader) {
	fmt.Printf("\n%s👉 Нажмите Enter, чтобы продолжить...%s", Cyan, Reset)
	reader.ReadString('\n')
}

func Run(s storage.Storage) {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		fmt.Printf("%s%s=== МЕНЕДЖЕР ЗАДАЧ ===%s\n\n", Bold, Cyan, Reset)
		fmt.Println("1. Показать задачи")
		fmt.Println("2. Добавить задачу")
		fmt.Println("3. Переключить статус")
		fmt.Println("4. Переименовать задачу")
		fmt.Println("5. Удалить задачу")
		fmt.Println("6. Статистика")
		fmt.Println("7. Выход")
		fmt.Println()

		choice, err := readInt(reader, "Выберите действие: ")
		if err != nil {
			PrintWarning("Пожалуйста, введите число от 1 до 7")
			waitEnter(reader)
			continue
		}

		switch choice {
		case 1:
			tasks := s.List()
			if len(tasks) == 0 {
				PrintInfo("Список задач пуст.")
				waitEnter(reader)
				continue
			}

			fmt.Println("\n1. Все задачи")
			fmt.Println("2. Только невыполненные")
			fmt.Println("3. Только выполненные")
			filterChoice, err := readInt(reader, "Выберите фильтр: ")
			if err != nil {
				PrintWarning("Некорректный выбор фильтра")
				waitEnter(reader)
				continue
			}

			fmt.Printf("\n%s--- Список задач ---%s\n", Bold, Reset)
			for _, t := range tasks {
				switch filterChoice {
				case 1:
					fmt.Println(t)
				case 2:
					if !t.Done {
						fmt.Println(t)
					}
				case 3:
					if t.Done {
						fmt.Println(t)
					}
				default:
					fmt.Println(t)
				}
			}
			waitEnter(reader)

		case 2:
			name := readLine(reader, "Введите название задачи: ")
			t, err := s.Add(name)
			if err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess(fmt.Sprintf("Задача #%d \"%s\" успешно добавлена!", t.ID, t.Name))
			waitEnter(reader)

		case 3:
			id, err := readInt(reader, "Введите номер задачи: ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			if err := s.Toggle(id); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Статус задачи изменён!")
			waitEnter(reader)

		case 4:
			id, err := readInt(reader, "Введите номер задачи: ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			name := readLine(reader, "Введите новое название: ")
			if err := s.Rename(id, name); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Задача переименована!")
			waitEnter(reader)

		case 5:
			id, err := readInt(reader, "Введите номер задачи для удаления: ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			if err := s.Delete(id); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Задача успешно удалена!")
			waitEnter(reader)

		case 6:
			total, completed, pending := s.Stats()
			fmt.Printf("\n%s--- Статистика задач ---%s\n", Bold, Reset)
			fmt.Printf("Всего задач:  %d\n", total)
			fmt.Printf("%sВыполнено:    %d%s\n", Green, completed, Reset)
			fmt.Printf("%sВ процессе:   %d%s\n", Yellow, pending, Reset)
			waitEnter(reader)

		case 7:
			PrintInfo("До встречи!")
			return

		default:
			PrintWarning("Неверный пункт меню, выберите от 1 до 7")
			waitEnter(reader)
		}
	}
}
