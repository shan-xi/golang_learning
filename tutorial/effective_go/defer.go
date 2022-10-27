package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// r, err := Contents("C:\\coding2\\golang_learning\\tutorial\\effective_go\\functions.go")
	// if err != nil {
	// 	fmt.Printf("error: %v", err)
	// }
	// fmt.Println(r)

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	b()
}

func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	return string(result), nil
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
