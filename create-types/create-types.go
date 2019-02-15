package main

import "fmt"

var a int

type hotdog int // create type // think of type just like var
var b hotdog // declare variable b of type hotdog

func main() {
	a = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 44
	fmt.Println(b)
	fmt.Printf("%T\n", b) // prints: main.hotdog // so 44 if of type hotdog from package main

	// a = b // cannot do because they are 2 different types
}
