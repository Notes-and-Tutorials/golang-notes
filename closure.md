# Closure
- want to enclose the scope of a variable to contain
- limited scope

```go
// package level scope
// scope is the entire package
var x int

func main() {
	fmt.Println(x)
	x++
	foo()
	fmt.Println(x)
}

func foo() {
	fmt.Println("hello")
	x++
}
```
narrow to main func
```go
func main() {
	var x int
	fmt.Println(x)
	x++

	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")
}
```

```go
func main() {
	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)

	{
	y := 42
	fmt.Println(y)
	}
	
	// fmt.Println(y) // can't print y here
}
```

```go

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println("hello", a())
	fmt.Println("hello", a())
	fmt.Println("hello", a())
	fmt.Println("bye", b())
	fmt.Println("bye", b())
}

func incrementor() func() int {
	var x int

	// no need to pass x into this function because it is in the scope of incrementor
	return func() int {
		x++
		return x
	}
}
```