package main

import "fmt"

type Wallet struct {
	Balance int
}

func main() {
	var w *Wallet
	if w == nil {
		fmt.Println("Кошелька не существует")
		return
	}
	if w != nil {
		fmt.Println(w.Balance)
	}
}
