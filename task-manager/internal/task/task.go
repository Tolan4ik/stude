package task

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func New(id int, name string) (Task, error) {
	if name == "" {
		return Task{}, ErrEmptyTitle // возвращаем нашу ошибку!
	}
	return Task{
		ID:   id,
		Name: name,
		Done: false,
	}, nil
}

func (t *Task) Toggle() {
	t.Done = !t.Done
}
func (t *Task) Rename(name string) error {
	if name == "" {
		return ErrEmptyTitle
	}
	t.Name = name
	return nil
}

func (t Task) String() string {
	status := "[❌ ]"
	if t.Done {
		status = "[✅ ]"
	}
	return fmt.Sprintf("%d. %s %s", t.ID, status, t.Name)
}
