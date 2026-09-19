package main

import (
	"log"
	"task-manager/internal/storage"
	"task-manager/internal/ui"
)

const storageFile = "tasks.json"

func main() {
	store, err := storage.NewFileStorage(storageFile)
	if err != nil {
		log.Fatalf("Ошибка инициализации хранилища: %v", err)
	}

	ui.Run(store)
}
