package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{1, 2, 3, 4}
	fmt.Println(Index(si, 4))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "b"))
}
