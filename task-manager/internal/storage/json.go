package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager/internal/task"
)

type FileStorage struct {
	filePath string
	tasks    []task.Task
}

func (f *FileStorage) save() error {

	data, err := json.MarshalIndent(f.tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования JSON: %w", err)
	}
	err = os.WriteFile(f.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка записи в файл:%w", err)
	}
	return nil
}

func (f *FileStorage) load() error {
	data, err := os.ReadFile(f.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			f.tasks = []task.Task{}
			return nil
		}
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}
	err = json.Unmarshal(data, &f.tasks)
	if err != nil {
		return fmt.Errorf("ошибка декодирования JSON: %w", err)
	}
	return nil
}

func NewFileStorage(filePath string) (*FileStorage, error) {

	storage := &FileStorage{
		filePath: filePath,
	}
	err := storage.load()
	if err != nil {
		return nil, err
	}
	return storage, nil

}

func (f *FileStorage) List() []task.Task {
	return f.tasks
}

func (f *FileStorage) Add(name string) (task.Task, error) {
	newID := len(f.tasks) + 1
	t, err := task.New(newID, name)
	if err != nil {
		return task.Task{}, err
	}
	f.tasks = append(f.tasks, t)
	if err := f.save(); err != nil {
		return task.Task{}, err
	}
	return t, nil
}

func (f *FileStorage) Get(id int) (task.Task, error) {
	for _, t := range f.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return task.Task{}, task.ErrTaskNotFound
}

func (f *FileStorage) Toggle(id int) error {
	for i := range f.tasks {
		if f.tasks[i].ID == id {
			f.tasks[i].Toggle() // меняем статус через метод структуры Task
			return f.save()     // сохраняем на диск
		}
	}
	return task.ErrTaskNotFound // если не нашли с таким id
}

func (f *FileStorage) Rename(id int, name string) error {

	for i := range f.tasks {
		if f.tasks[i].ID == id {
			if err := f.tasks[i].Rename(name); err != nil {
				return err
			}
			return f.save()
		}

	}
	return task.ErrTaskNotFound
}

func (f *FileStorage) Delete(id int) error {

	foundIndex := -1
	for i, t := range f.tasks {
		if t.ID == id {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return task.ErrTaskNotFound
	}

	f.tasks = append(f.tasks[:foundIndex], f.tasks[foundIndex+1:]...)
	for i := foundIndex; i < len(f.tasks); i++ {
		f.tasks[i].ID = i + 1
	}
	return f.save()
}

func (f *FileStorage) Stats() (total, completed, pending int) {
	total = len(f.tasks)
	for _, t := range f.tasks {
		if t.Done {
			completed++
		} else {
			pending++
		}
	}
	return total, completed, pending
}
