package main

import "fmt"

var a int

type hotdog int // create type
var b hotdog // declare variable b of type hotdog

func main() {
	a = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 44
	fmt.Println(b)
	fmt.Printf("%T\n", b) // prints: main.hotdog // so 44 if of type hotdog from package main

	a = int(b) // Type conversion (called casting in other languages)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
