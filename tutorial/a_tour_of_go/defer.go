package main

import "fmt"

func main() {
	defer fmt.Println("w")
	fmt.Println("H")
}
