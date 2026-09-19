package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-manager/internal/storage"
	"task-manager/internal/task"
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

func printTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		PrintInfo("Список задач пуст.")
		return
	}
	fmt.Printf("\n%s--- Список задач ---%s\n", Bold, Reset)
	for _, t := range tasks {
		fmt.Println(t)
	}
}

func promptShowTasks(reader *bufio.Reader, s storage.Storage) {
	fmt.Printf("\n%s👉 [1] Показать задачи | [Enter] Главное меню: %s", Cyan, Reset)
	choice := readLine(reader, "")
	if choice == "1" {
		printTasks(s.List())
		waitEnter(reader)
	}
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

		choice, err := readInt(reader, "Выберите действие (1-7): ")
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

			for {
				clearScreen()
				fmt.Printf("%s%s=== ПРОСМОТР ЗАДАЧ ===%s\n\n", Bold, Cyan, Reset)
				fmt.Println("1. Все задачи")
				fmt.Println("2. Только невыполненные")
				fmt.Println("3. Только выполненные")
				fmt.Println("0. Назад в главное меню")
				fmt.Println()

				filterChoice, err := readInt(reader, "Выберите фильтр (0-3): ")
				if err != nil {
					PrintWarning("Некорректный выбор фильтра")
					waitEnter(reader)
					continue
				}

				if filterChoice == 0 {
					break // уходим в главное меню
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

				// ВОТ ЗДЕСЬ НАШ ВЫБОР НАВИГАЦИИ:
				fmt.Printf("\n%s👉 [1] Выбрать другой фильтр | [Enter] Главное меню: %s", Cyan, Reset)
				navChoice := readLine(reader, "")
				if navChoice != "1" {
					break // если не 1, выходим в главное меню
				}
			}

		case 2:
			name := readLine(reader, "Введите название задачи: ")
			t, err := s.Add(name)
			if err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess(fmt.Sprintf("Задача #%d \"%s\" успешно добавлена!", t.ID, t.Name))
			promptShowTasks(reader, s)

		case 3:
			tasks := s.List()
			if len(tasks) == 0 {
				PrintInfo("Список задач пуст.")
				waitEnter(reader)
				continue
			}

			printTasks(tasks) // 👈 ВОТ ОНО: выводим задачи перед глазами!
			id, err := readInt(reader, "\nВведите номер задачи для изменения статуса (или 0 для отмены): ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			if id == 0 {
				continue
			}
			if err := s.Toggle(id); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Статус задачи изменён!")
			promptShowTasks(reader, s)

		case 4:
			tasks := s.List()
			if len(tasks) == 0 {
				PrintInfo("Список задач пуст.")
				waitEnter(reader)
				continue
			}

			printTasks(tasks)
			id, err := readInt(reader, "\nВведите номер задачи для изменения названия (или 0 для отмены): ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			if id == 0 {
				continue
			}
			name := readLine(reader, "Введите новое название: ")
			if err := s.Rename(id, name); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Задача переименована!")
			promptShowTasks(reader, s)

		case 5:
			tasks := s.List()
			if len(tasks) == 0 {
				PrintInfo("Список задач пуст.")
				waitEnter(reader)
				continue
			}

			printTasks(tasks)
			id, err := readInt(reader, "\nВведите номер задачи для удаления (или 0 для отмены): ")
			if err != nil {
				PrintWarning("Некорректный номер задачи")
				waitEnter(reader)
				continue
			}
			if id == 0 {
				continue
			}
			if err := s.Delete(id); err != nil {
				PrintError(err)
				waitEnter(reader)
				continue
			}
			PrintSuccess("Задача успешно удалена!")
			promptShowTasks(reader, s)

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
