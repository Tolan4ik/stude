package main

import "fmt"

func addTen(p *int) {

	*p = *p + 10
}

func main() {

	val := 40
	addTen(&val)
	fmt.Println(val)
}
