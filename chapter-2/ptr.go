package main

import "fmt"

func main() {
	v := 1
	p := &v

	fmt.Println("v: ", v)
	fmt.Println("*p: ", *p)

	*p++
	fmt.Println("v: ", v)
	fmt.Println("*p: ", *p)

	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)

	var p2 *int
	fmt.Println(p2)

	fmt.Println("-----------------")
	p3 := f()
	fmt.Println(*p3)
	fmt.Println(f() == f())

	v2 := 1
	incr(&v2)
	fmt.Println(incr(&v2))
}

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}


