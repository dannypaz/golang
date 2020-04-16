package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		lines := make(map[string]int)
		countLines(os.Stdin, lines)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			lines := make(map[string]int)
			countLines(f, lines)
			counts[arg] = lines
		}
	}

	for file, count := range counts {
		for line, n := range count {
			if n < 1 {
				continue
			}
			fmt.Printf("%s\t%d\t%s\n", file, n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
