// Dup1 prints the text of each line that appears more than
// once in the stdin, preceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Use ctrl+D to break out of this command
	for input.Scan() {
		counts[input.Text()]++
	}

	// Note: Ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
