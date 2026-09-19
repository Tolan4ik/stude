package task

import "errors"

var (
	ErrEmptyTitle   = errors.New("название задачи не может быть пустым")
	ErrTaskNotFound = errors.New("Задача с таким ID не найдена")
)
