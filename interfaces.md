# Interfaces
- allows for polymorphism 
- allow us to define behavior
- allow a value to be of more than 1 type

> general: a value can be of more than 1 type

```go
// Type person has the method speak attached to it
type person struct {
	first string
	last  string
	age   int
}

// Type secretAgent has the method speak attached to it
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("i am", s.first, s.last, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("i am", p.first, p.last, "- the person speak")
}

// Interface
// `keyword identifier type`
// any type that has the method speak is also of type human
type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	// so this value is of type secretAgent
	// but because it has the method speak() attached to it, it is also of type human
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			32,
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Noe",
	}

	fmt.Println(sa1)
	sa1.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(p1)

} // end main

```