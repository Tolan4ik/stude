package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func move(p Point) {

	p.X += 10
}

func movePtr(p *Point) {

	p.X += 10
}

func main() {
	p := Point{X: 0, Y: 0}
	move(p)
	fmt.Println(p.X)

	movePtr(&p)
	fmt.Println(p.X)
}
