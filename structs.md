# Structs
- data structure
- allows us to compose together values of different types
- an aggregate data type
- complex data types

- similar to class or object 
> "We have created a value of type ...

```go
// type <idenfifier(name)> struct
// struct is the underlying type
// composite data structure
type person struct {
	// fields
	first   string
	last    string
	age     int
	parents []string
}

func main() {
	// create values of type person
	p1 := person{
		first:   "hailee",
		last:    "miu",
		age:     27,
		parents: []string{"Shelly", "John"}, // need trailing commas
	}
	p2 := person{
		first:   "nikko",
		last:    "miu",
		age:     26,
		parents: []string{"Julie"}, // need trailing commas
	}
	fmt.Println(p1)
	fmt.Println(p2)

	// can use dot notation
	fmt.Println(p1.first, p1.last, p1.age, p1.parents)
	fmt.Println(p2.first, p2.last, p2.age, p2.parents)
}
```

## Embedded Structs
-  

```go
type person struct {
	first string
	last  string
	age   int
}

// A secretAgent is everything that a person is, but with a ltk (license to kill)
// Embed `person` into `secretAgent`
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "nikko",
		last:  "miu",
		age:   26,
	}
	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk) // inner part "gets promoted" (unless collision)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age, sa1.ltk) // but still can do like this if there is name collision
	fmt.Println(p2.first, p2.last, p2.age)
}
```

## Anonymous structs
- do if don't want code pollution
- nothing extraneous 
```go
func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "J",
		last:  "B ",
		age:   32,
	}
	fmt.Println(p1)
}
```