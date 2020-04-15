package main

import (
	"fmt"
)

func main() {
	var	arr [5]int

	for i := range arr {
		fmt.Println("Hello!", i)
	}
}
