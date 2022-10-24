package main

import (
	"fmt"
)

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {

	e1 := List[int]{}
	e1.val = 1
	e2 := List[int]{}
	e2.val = 2
	e1.next = &e2

	fmt.Println(e1.val)
	fmt.Println((e1.next).val)
}
