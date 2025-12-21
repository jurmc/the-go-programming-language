package main

import "fmt"

func main() {
	s := "Hello, \x01\x30\x00\x01\x30\x01"
	fmt.Println(s)
	fmt.Println(len(s))
}
