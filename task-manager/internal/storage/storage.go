package storage

import "task-manager/internal/task"

type Storage interface {
	Add(name string) (task.Task, error)
	Get(id int) (task.Task, error)
	List() []task.Task
	Toggle(id int) error
	Rename(id int, name string) error
	Delete(id int) error
	Stats() (total, completed, pending int)
}
