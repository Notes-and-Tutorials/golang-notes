# Recursion
- anything you do with recursion, you could also do with loops
- when a function calls itself
- always need some way for it to stop

```go 
func main() {
	n := factorial(4)

	fmt.Println(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
```

### using loops instead
```go
func main() {
	n := factorial(4)
	fmt.Println(n)
	
	l := factorial(4)
	fmt.Println(l)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n>0; n-- { // left init empty because didn't need
		total *= n
	}
	return total
}
```