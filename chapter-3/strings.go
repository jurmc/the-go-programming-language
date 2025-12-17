package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println(s[5:8])
	fmt.Println("goodbye " + s[7:])

	s2 := "left foot"
	t := s2
	s2 += ", right foot"
	fmt.Println(s2)
	fmt.Println(t)
}
