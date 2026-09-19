package main

type Wallet struct {
	Balance int
}

func StayOnStack() Wallet {
	w := Wallet{Balance: 50}
	return w
}

func EscapeToHeap() *Wallet {
	w := Wallet{Balance: 100}
	return &w
}

func main() {
	_ = StayOnStack()
	_ = EscapeToHeap()
}
