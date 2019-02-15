package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello", 34, false)
	fmt.Println(n)
	fmt.Println(err)
	// Prints:
		// 	Hello 34 false
		// 15 <----number of bytes
		// <nil>
}

/*  
Throw away a return because you can not ignore a variable in Go

func main() {
	n, _ := fmt.Println("Hello", 34, false)
	fmt.Println(n)
}
*/