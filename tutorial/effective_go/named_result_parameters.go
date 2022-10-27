package main

import "fmt"

func a() (g int) {
	g = 1
	return
}

func main() {
	var b int
	b = a()
	fmt.Println(b)
}
