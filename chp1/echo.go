package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Name of the program: ", os.Args[0])
	fmt.Println("Start time:", time.Now())
	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Println("idx:", i, "arg:", arg)
	}
	end := time.Now()
	fmt.Println("Time duration:", end.Sub(start))

	start = time.Now()
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
	end = time.Now()
	fmt.Println("Time duration:", end.Sub(start))
}
