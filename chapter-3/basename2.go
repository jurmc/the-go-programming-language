package main

import "fmt"
import "strings"

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		s = s[:dot]
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
