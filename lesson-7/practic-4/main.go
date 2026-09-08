package main

import (
	"fmt"
	"unsafe"
)

var p1 *int
var p2 *string
var p3 *struct{}

func main() {

	fmt.Println(unsafe.Sizeof(p1))
	fmt.Println(unsafe.Sizeof(p2))
	fmt.Println(unsafe.Sizeof(p3))
}
