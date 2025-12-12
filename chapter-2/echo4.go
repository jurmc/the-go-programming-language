package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
)

func main() {
	fmt.Println("os.Args-----------")
	for _, a := range os.Args {
		fmt.Println(a)
	}



	fmt.Println("flag.Args-----------")
	n := flag.Bool("n", false, "ommit trailing newline")
	sep := flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Printf("*n: %v, *sep: %v\n", *n, *sep)
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	
}
