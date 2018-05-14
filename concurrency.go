package main

import (
	"fmt"
	"time"
)

// The sync package provides useful primitives,
// although you won't need them much in Go as there are other primitives.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// goroutine
	go say("world")
	say("hello")
}
