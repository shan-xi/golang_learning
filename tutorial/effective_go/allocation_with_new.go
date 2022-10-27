package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	p := new(SyncedBuffer)
	var v SyncedBuffer

	fmt.Printf("%T\n", p)
	fmt.Printf("%T", v)
}
