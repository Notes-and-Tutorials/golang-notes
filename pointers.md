# Pointers
- pointing to an address in memory where a value is stored
- `&` gives you the address where the value is stored
- `*` when talking about the type

> everything in Go is `pass by value`
## When to use pointers
- if have large chunk of data that you do not want to pass around
  - save on performance
- need to change something that is at a certain location

```go
func main() {
	age := 27
	fmt.Println(age)  // 42
	fmt.Println(&age) // 0x416020

	fmt.Printf("%T\n", age)  // int
	fmt.Printf("%T\n", &age) // *int // * here is part of the type (pointer to an int) (where an int is stored
	
	b := &age
	fmt.Println(b) // 0x416020
	fmt.Println(*b) // 27 // dereferencing an address // gives value stored at that address
}
```
Invalid code: `var b int = &a`

### step 1 
```go
func main() {
	x := 0 
	foo(x)
	fmt.Println(x)
}

func foo(y int){
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}
```
### step 2
Address stays the same, but the value changes
```go
func main() {
	x := 0
	fmt.Println("&x before", &x)
	fmt.Println("x before", x)
	
	foo(&x)
	
	fmt.Println("&x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("*y before", *y)
	
	*y = 43
	
	fmt.Println("y after", y)
	fmt.Println("*y after", *y)
}
```
## Method Sets
- the methods attached to a type are know as its methods set
- 
"
a NON-POINTER RECEIVER
works with values that are POINTERS or NON-POINTERS.
a POINTER RECEIVER
only works with values that are POINTERS.

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

"

> If you want to point to something, use the `&`. If you want to read through a memory address, use the `*`.

