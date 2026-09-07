package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	id   int
	name string
	done bool
}

func (t *Task) Complete() {
	t.done = true
}

func (t *Task) SetID(id int) {
	t.id = id
}

func NewTask(id int, name string) (Task, error) {
	if name == "" {
		return Task{}, fmt.Errorf("название не может быть пустым")
	}

	return Task{
		id:   id,
		name: name,
		done: false,
	}, nil
}

func FindTaskIndex(tasks []Task, id int) (int, error) {
	for i, task := range tasks {
		if task.id == id {
			return i, nil
		}
	}

	return -1, fmt.Errorf("задача с таким номером не найдена")

}

func GetStats(tasks []Task) (total, completed, pending int) {
	total = len(tasks)
	for _, task := range tasks {
		if task.IsDone() {
			completed++
		} else {
			pending++
		}

	}
	return

}

func FilterTasks(tasks []Task, predicate func(Task) bool) []Task {
	var result []Task
	for _, task := range tasks {
		if predicate(task) {
			result = append(result, task)
		}

	}
	return result
}

func (t Task) Print() {
	fmt.Println("Номер задачи:", t.id)
	fmt.Println("Задача:", t.name)

	if t.IsDone() {
		fmt.Println("Выполнено")
	} else {
		fmt.Println("Не выполнено")
	}

	fmt.Println()
}

func (t Task) IsDone() bool {
	return t.done
}

func (t *Task) Toggle() {
	t.done = !t.done
}

func (t *Task) Rename(name string) {
	t.name = name
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var tasks []Task

	for {
		var choice int
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Изменить статус")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Переименовать задачу")
		fmt.Println("6. Показать статистику")
		fmt.Println("7. Выход")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {

		case 1:
			var name string

			fmt.Print("Введите название задачи:")

			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			task, err := NewTask(len(tasks)+1, name)
			if err != nil {
				fmt.Println("Поле не может быть пустым", err)
				break
			}
			fmt.Println("Задача добавлена!")
			tasks = append(tasks, task)

		case 2:

			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
				fmt.Println()
				break
			}

			var filterChoice int
			fmt.Println("1. Все задачи")
			fmt.Println("2. Только невыполненные")
			fmt.Println("3. Только выполненные")
			fmt.Print("Выберите фильтр: ")
			fmt.Scan(&filterChoice)

			var filtered []Task

			switch filterChoice {

			case 1:
				filtered = FilterTasks(tasks, func(t Task) bool { return true })
			case 2:
				filtered = FilterTasks(tasks, func(t Task) bool { return !t.IsDone() })
			case 3:
				filtered = FilterTasks(tasks, func(t Task) bool { return t.IsDone() })
			default:
				fmt.Println("Неверный пункт")

			}

			for _, task := range filtered {
				task.Print()
			}

		case 3:
			var id int
			fmt.Println("Введите номер задачи")
			fmt.Scan(&id)
			i, err := FindTaskIndex(tasks, id)
			if err != nil {
				fmt.Println(err)
				break
			}
			tasks[i].Toggle()
			fmt.Println("Статус изменён!")

		case 4:
			var id int
			fmt.Println("Введите номер задачи для удаления:")
			fmt.Scan(&id)

			i, err := FindTaskIndex(tasks, id)
			if err != nil {
				fmt.Println(err)
				break
			}

			tasks = append(tasks[:i], tasks[i+1:]...)
			for i := range tasks {
				tasks[i].SetID(i + 1)
			}

			fmt.Println("Задача удалена!")

		case 5:
			var id int

			fmt.Println("Введите номер задачи для переименования:")
			fmt.Scan(&id)
			reader.ReadString('\n')
			i, err := FindTaskIndex(tasks, id)
			if err != nil {
				fmt.Println(err)
				break
			}

			fmt.Println("Введите новое название:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			tasks[i].Rename(name)

			fmt.Println("Название изменено!")
		case 6:
			total, completed, pending := GetStats(tasks)
			fmt.Printf("Всего:%d", total)
			fmt.Printf("\nВыполненых:%d", completed)
			fmt.Printf("\nВ процессе:%d\n", pending)

		case 7:
			fmt.Println("Выход из программы...")
			return
		}

	}

}
