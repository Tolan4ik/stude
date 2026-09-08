package main

import (
	"encoding/json"
	"fmt"
)

type Transaction struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"description,omitempty"`
	SecretNote  string `json:"-"`
}

func main() {

	tx := Transaction{}
	tx.ID = "tx-1001"
	tx.Amount = 5000
	tx.Currency = "KZT"
	tx.Description = ""
	tx.SecretNote = "проверить паспорт клиента"

	data, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(string(data))
}
