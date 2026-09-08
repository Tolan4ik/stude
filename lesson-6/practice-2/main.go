package main

import (
	"fmt"
)

type Contact struct {
	Email string
}

type Employee struct {
	Contact
	Name string
}

func main() {

	employ := Employee{Name: "Толян", Contact: Contact{Email: "tola@mail.ru"}}
	fmt.Println(employ.Email)
}
