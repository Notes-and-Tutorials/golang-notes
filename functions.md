# Functions
- functions are all about being modular (small chunks)
  - can use functions and packages
- everything in Go is PASS BY VALUE
> Syntax: `func (r receiver) identifier(parameters) (returns(s)) {...}`

```go
func main() { // what goes in here are arguments
	foo()
	bar("Haile")
	fb := foobar("NIkko")
	fmt.Println(fb)
	a, b := taco("Bandit", "miu")
	fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) { // here in () are parameters
	fmt.Println("Hello,", s)
}

func foobar(x string) string {
	return fmt.Sprint("Hello from foobar, ", x)

}

func taco(x string, y string) (string, bool) {
	c := fmt.Sprint(x, y, `"says"hello"`)
	d := false
	return c, d

}
```

##Variadic Parameter
`...` operator
> can actually pass in nothing  
> must be the last parameter in the list

```go
func main() {
	foo(2, 4, 64, 478, 34) // turns into a slice of int to be able to store
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
```

```go
func main() {
	x := sum(2, 4, 64, 478, 34) // turns into a slice of int to be able to store
	fmt.Println("TOTAL:", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	// var sum int
	for i, v := range x {
		sum += v
		fmt.Println("for item in index", i, "we are adding", v, "to total", sum)

	}
	return sum
}
```

## Unfurling a slice
- HIS WORDS

```go
func main() {
	xi := []int{2, 4, 64, 478, 34}
	x := sum(xi...) // allows the ints to be passed in because sum parameters are variatic parameters
	fmt.Println("TOTAL:", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	// var sum int
	for i, v := range x {
		sum += v
		fmt.Println("for item in index", i, "we are adding", v, "to total", sum)

	}
	return sum	
}
```

## Defer
- defer execution of a function until wherever its being called comes to an end

- could use to organize code so that close file is right after open file, but just defered
```go
func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
```
