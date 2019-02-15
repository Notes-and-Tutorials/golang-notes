# Grouping Data  
Or called aggregate data

> Syntax difference between array and slice, slice does not have the length inside of the []

## Arrays
- Arrays are not usually used in Go
> Use slices instead

- Must be a single type
- Length is part of its type

```go
func main() {
	var x [5]int
	fmt.Println(x)
	x[3] = 12
	fmt.Println(x)
	fmt.Println(len(x)) // length
}
```

## Slice
Composite Literal: `x := type{value}`

allows you to group together values of the same type

```go
func main() {
	x := []int{12, 13, 14}
	fmt.Println(x)
}
```

### Loop over values

```go
func main() {
	x := []int{10, 11, 12}
	fmt.Println(len(x))
	fmt.Println(x[1]) // access by index position

	for i, v := range x {
		fmt.Println(i, v)
	}
}
```
### Slicing a slice
- how you would delete from a slice

```go
func main() {
	x := []int{10, 11, 12, 13}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
}
```

### Append to a slice

```go
func main() {
	x := []int{10, 11, 12, 13}
	fmt.Println(x)
	x = append(x, 14, 15)
	fmt.Println(x)

	y := []int{1, 2, 3}
	x = append(x, y...)
	fmt.Println(x)
}
```
### Delete from slice

```go
func main() {
	x := []int{10, 11, 12, 13}
	fmt.Println(x)
	x = append(x, 14, 15)
	fmt.Println(x)

	y := []int{1, 2, 3}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:5], x[7:]...) // remove 15 and 1 // ... on right hand side is called unfurling 
	fmt.Println(x)
}
```
### slice - make
- a different way to create a slice
> a slice is built on top of an array (in general)
- use make when you know you have a fixed number of items

```go
func main() {
	// slice, length, capacity of the underlying array
	x := make([]int, 11, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// assign specifics
	x[0] = 12
	x[9] = 89
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// append to increase length
	x = append(x, 111)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// doubles the size of original capacity if goes over
	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
```
### Multi-dimensional slice
- like a spread sheet
```go
func main() {
	jb := []string{"James", "Bond", "Cholocate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Peach"}
	fmt.Println(mp)

	// a slice of a slice of string
	// multidimensional slice
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
```

#### Range over  slice of a slice of a string then range over the records, then range over the data
```go
func main() {
	b := []string{"Bandit", "kitty", "meows", "cuddles"}
	t := []string{"Taurus", "kitty", "sleeps", "cuddles"}

	fmt.Println(b)
	fmt.Println(t)

	bt := [][]string{b, t}
	fmt.Println(bt)

	for i, x := range bt {
		fmt.Println("record: ", i)
		for j, val := range x {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

}
```