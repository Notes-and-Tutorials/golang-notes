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
## Anonymous Function
- anonymous self executing functions
```go
func main() {
	foo()

	// anonymous
	// Don't forget the parenthesis at the end
	func() {
		fmt.Println("anonymous fun ran")
	}()

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42) // pass in arguments here
}

// not anonymous
func foo() {
	fmt.Println("foo ran")
}
```

## Func expression
- function assigned to a variable (assign a function to a variable)
- functions are first class citizens (it is a type, just like all other types)

```go
func main() {
	a := func() {
		fmt.Println("my first func expression")
	}
	a()

	b := func(x int) {
		fmt.Println("answer", x)
	}
	b(6)
}
```

## Return a function from a function 
(function that returns a function)

#### remember returning a string
```go
func main() {
	s1 := foo()
	fmt.Println(s1)
}
func foo() string {
	s := "hello world"
	return s
}
```
#### both examples
```go
func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()
	fmt.Printf("%T\n", x)
	
	i := x()
	fmt.Println(i)
}
func foo() string {
	// s := "hello world"
	// return s
	// OR
	return "hello world"
}

// the return is a function that returns and int
func bar() func() int {
	// techincally an anonymous function
	return func() int {
		return 6
	}
}
```
#### function that returns an int
```go
func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func bar() func() int {
	return func() int {
		return 6
	}
}

```
step cleaner
```go
func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	fmt.Println(x())
}

func bar() func() int {
	return func() int {
		return 6
	}
}
```
cleanest
```go
func main() {
	fmt.Println(bar()())
}

func bar() func() int {
	return func() int {
		return 6
	}
}
```
## Call back
- passing a func as an argument
- technically functional programming

### Example sum
```go
func main() {
	ii := []int{12, 14, 445, 6, 4567}
	s := sum(ii...)
	fmt.Println(s)
}

func sum(xi ...int) int {
	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
```

### Example sum even and odd
```go
func main() {
	ii := []int{12, 14, 445, 6, 4567}
	s := sum(ii...)
	fmt.Println("all numbers", s)

	s2 := even(sum, ii...)
	fmt.Println("even numbers", s2)

	s3 := odd(sum, ii...)
	fmt.Println("odd numbers", s3)
}

func sum(xi ...int) int {
	total := 0r
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(a ...int) int, b ...int) int {
	var c []int
	for _, d := range b {
		if d%2 != 0 {
			c = append(c, d)
		}
	}
	return f(c...)
}
```