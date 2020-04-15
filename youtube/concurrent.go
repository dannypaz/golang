package main

import (
	"fmt"
	"time"
)

// Go... you send channels between go routines using channels
// channels are types
// buffered channels will not block...
// 

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
