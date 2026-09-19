package main

import "fmt"

type Client struct {
	name string
	next *Client
}

func printQueue(head *Client) {

	for c := head; c != nil; c = c.next {
		fmt.Print(c.name, " ➡️  ")
	}
	fmt.Println("Конец очереди")
}

func main() {

	marat := &Client{name: "Марат"}
	ivan := &Client{name: "Иван", next: marat}
	head := ivan
	fmt.Println("Обычная очередь:")
	printQueue(head)

	vip := &Client{name: "VIP-Алибек", next: head}
	head = vip
	fmt.Println("\nОчередь после прихода VIP:")
	printQueue(head)
}
