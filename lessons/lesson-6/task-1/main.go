package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	done bool
	id   int64
	say  bool
}

type Good struct {
	id   int64
	done bool
	say  bool
}

func main() {

	sumBad := unsafe.Sizeof(Bad{})
	sumGood := unsafe.Sizeof(Good{})
	fmt.Println(sumBad)
	fmt.Println(sumGood)

}
