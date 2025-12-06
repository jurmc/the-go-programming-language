// Echo prints its command-line args
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	fmt.Println("echo1")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println("echo2")
	for i, arg := range os.Args[1:] {
		fmt.Println(i, ": ", arg)
	}
}

func echo3() {
	fmt.Println("echo3")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func echo4() {
	fmt.Println("echo4")
	fmt.Println(os.Args[0:], " ")
}

func main() {
	echo1()
	echo2()
	echo3()
	echo4()
}
