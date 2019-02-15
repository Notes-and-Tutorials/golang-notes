package main

import "fmt"
// iota predeclared identifier
// automatically increment by 1

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// const (
// 	a = iota + 1
// 	b = iota
// 	c = iota
// )

const (
	a = iota
	b
	c
)

const (
	e = iota
	f
	g
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(g)
	fmt.Printf("%T\n", g)
}


