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

func NewTask(id int, name string) Task {
	return Task{
		id:   id,
		name: name,
		done: false,
	}
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
		fmt.Println("6. Выход")
		fmt.Print("Выберите действие: ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {

		case 1:
			var name string

			fmt.Print("Введите название задачи:")

			name, _ = reader.ReadString('\n')
			name = strings.TrimSpace(name)

			task := NewTask(len(tasks)+1, name)
			tasks = append(tasks, task)
			fmt.Println("Задача добавлена!")
		case 2:

			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
				fmt.Println()
				break
			}
			for _, task := range tasks {
				task.Print()
			}

		case 3:
			var id int
			fmt.Println("Введите номер задачи")
			fmt.Scan(&id)
			found := false
			for i := range tasks {
				if tasks[i].id == id {
					tasks[i].Toggle()
					found = true
					fmt.Println("Статус изменён!")
					break
				}

			}
			if !found {
				fmt.Println("Такой задачи нет")
			}

		case 4:
			var id int
			fmt.Println("Введите номер задачи для удаления:")
			fmt.Scan(&id)
			found := false
			for i := range tasks {
				if tasks[i].id == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					for i := range tasks {
						tasks[i].SetID(i + 1)
					}
					found = true
					fmt.Println("Задача удалена!")
					break
				}

			}
			if !found {
				fmt.Println("Такой задачи нет")
			}

		case 5:
			var id int

			fmt.Println("Введите номер задачи для переименования:")
			fmt.Scan(&id)
			reader.ReadString('\n')

			found := false

			for i := range tasks {
				if tasks[i].id == id {

					fmt.Println("Введите новое название:")
					name, _ := reader.ReadString('\n')
					name = strings.TrimSpace(name)

					tasks[i].Rename(name)

					found = true
					fmt.Println("Название изменено!")
					break
				}
			}

			if !found {
				fmt.Println("Такой задачи нет")
			}

		case 6:
			fmt.Println("Выход из программы...")
			return
		}

	}

}
