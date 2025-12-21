package main

import "fmt"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}

	return s
}

func main() {
	paths := []string{
		"/a/b/c",
		"b",
		"/xyz/c.go"}
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		fmt.Printf("path: %s, basename %s\n", path, basename(path))
	}
}
