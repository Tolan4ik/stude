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

func handleTaskAction(reader *bufio.Reader, s storage.Storage, id int) {
	t, err := s.Get(id)
	if err != nil {
		PrintError(err)
		waitEnter(reader)
		return
	}

	for {
		clearScreen()
		fmt.Printf("%s%s=== ДЕЙСТВИЯ НАД ЗАДАЧЕЙ ===%s\n\n", Bold, Cyan, Reset)
		fmt.Printf("Выбрана задача: %s\n\n", t)
		fmt.Println("1. Переключить статус")
		fmt.Println("2. Переименовать")
		fmt.Println("3. Удалить")
		fmt.Println("0. Назад к списку")
		fmt.Println()

		action, err := readInt(reader, "Выберите действие (0-3): ")
		if err != nil {
			PrintWarning("Пожалуйста, введите число от 0 до 3")
			waitEnter(reader)
			continue
		}

		switch action {
		case 1:
			if err := s.Toggle(id); err != nil {
				PrintError(err)
			} else {
				PrintSuccess("Статус изменён!")
			}
			waitEnter(reader)
			return // сразу возвращаемся к списку задач

		case 2:
			name := readLine(reader, "Введите новое название (или Enter для отмены): ")
			if name == "" {
				PrintInfo("Переименование отменено.")
				waitEnter(reader)
				return // просто возвращаемся к списку задач
			}
			if err := s.Rename(id, name); err != nil {
				PrintError(err)
			} else {
				PrintSuccess("Задача переименована!")
			}
			waitEnter(reader)
			return // сразу возвращаемся к списку задач

		case 3:
			if err := s.Delete(id); err != nil {
				PrintError(err)
			} else {
				PrintSuccess("Задача успешно удалена!")
			}
			waitEnter(reader)
			return // сразу возвращаемся к списку задач

		case 0:
			return // просто назад к списку задач

		default:
			PrintWarning("Неверный пункт, выберите от 0 до 3")
			waitEnter(reader)
		}
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
			filterChoice := 1 // по умолчанию показываем все задачи

			for {
				tasks := s.List()
				if len(tasks) == 0 {
					PrintInfo("Список задач пуст.")
					waitEnter(reader)
					break
				}

				clearScreen()
				fmt.Printf("%s%s=== ВАШИ ЗАДАЧИ ===%s\n\n", Bold, Cyan, Reset)

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

				fmt.Printf("\n%s👉 Номер задачи для действий | [F] Фильтры | [Enter] Меню: %s", Cyan, Reset)
				input := readLine(reader, "")

				// Если просто Enter — уходим в главное меню
				if input == "" {
					break
				}

				// Если ввели F — меняем фильтр
				if strings.ToLower(input) == "f" {
					fmt.Println("\n1. Все задачи")
					fmt.Println("2. Только невыполненные")
					fmt.Println("3. Только выполненные")
					fChoice, err := readInt(reader, "Выберите фильтр (1-3): ")
					if err == nil && fChoice >= 1 && fChoice <= 3 {
						filterChoice = fChoice
					}
					continue
				}

				// Если ввели число — это номер задачи!
				taskID, err := strconv.Atoi(input)
				if err != nil {
					PrintWarning("Неверный ввод (введите номер задачи, F или Enter)")
					waitEnter(reader)
					continue
				}

				// Запускаем действия над выбранной задачей!
				handleTaskAction(reader, s, taskID)
			}

		case 2:
			name := readLine(reader, "Введите название задачи (или Enter для отмены): ")
			if name == "" {
				PrintInfo("Добавление отменено.")
				waitEnter(reader)
				continue
			}

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
			name := readLine(reader, "Введите новое название (или Enter для отмены): ")
			if name == "" {
				PrintInfo("Переименование отменено.")
				waitEnter(reader)
				continue // просто возвращаемся в меню
			}
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
