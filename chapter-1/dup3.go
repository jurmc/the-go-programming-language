package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files_c := make(map[string]bool)
	files := os.Args[1:]

	for _, filename := range files {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, filename, files_c)
		f.Close()
	}

	for file, contains := range files_c {
		if contains {
			fmt.Printf("file: %s\n", file)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename string, files_c map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			files_c[filename] = true
		}
	}
	// NOTE: ignoring potential error from input.Err()
}
