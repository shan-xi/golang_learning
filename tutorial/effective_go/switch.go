package main

import "fmt"

func main() {
	src := []int{100, 200, 1, 400}
	sizeOne := 10
	sizeTwo := 20
	validateOnly := true
	sum := 0
Loop:
	for n := 0; n < len(src); n++ {
		fmt.Println("n=%d", n)

		switch {
		case src[n] < sizeOne:
			if validateOnly {
				fmt.Println("break")
				break
			}

		case src[n] < sizeTwo:
			fmt.Println("loop")
			break Loop
		}
		sum++
	}

	fmt.Println(sum)
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}
