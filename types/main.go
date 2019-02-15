// types are what make Go a static language (JS is dynamic)
package main

import "fmt"

// All of these are at package scope
var y = 42
var z string = "words"
// OR // var z = "words"
var a = 
`variable 
with 
"quotes" and spacing`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// z = 43 // cannot reassign to a different type // ERROR: cannot use 43 (type int) as type string in assignment

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
