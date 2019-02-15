package main

import "fmt"

// Untyped constants
// const a = 43
// const b = 44.345
// const c = "me"

const (
	a = 43
	b = 44.345
	c = "me"
)

// Typed constants
// const (
// 	a int = 43
// 	b float64 = 44.345
// 	c string = "me"
// )

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)
}
 