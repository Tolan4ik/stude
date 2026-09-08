package main

import "fmt"

func inc(p *int) {
	*p = *p + 1
}

func main() {
	x := 3
	inc(&x)
	fmt.Println(x)
}
