# Interfaces
- allows for polymorphism 
- allow us to define behavior
- allow a value to be of more than 1 type
- `polymorphism (multiple versions of a code block/function (of same name) that take in different types of arguments. So version that gets used depends on what type it is given)`

> general: a value can be of more than 1 type

> "Interface says, hey baby, if you've go this method, then you're my type."

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

// Method
func (s secretAgent) speak() {
	fmt.Println("i am", s.first, s.last, "- the secretAgent speak")
}

// Method
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
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		"Dr.",
		"Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	p1.speak()
	
	// So now this function can take in many types
	bar(sa1)
	bar(sa2)
	bar(p1)

} // end main

```

#### my example
```go
package main

import (
	"fmt"
)

type pet struct {
	name   string
	mammal bool
	legs   int
}

type cat struct {
	pet
	short_hair bool
}

type bird struct {
	pet
	tropical bool
}

// Method
func (c cat) speak() {
	fmt.Println("Meow-", c.name)
}

// Method
func (b bird) speak() {
	fmt.Println("Chirp-", b.name)
}

// Method
func (b bird) fly() {
	fmt.Println("flying-", b.name)
}

// Any type that has the method speak on it, is also type animal
type animal interface {
	speak()
	fly()
}

// both birds and cats are animals because they can both speak
func powerCall(a animal) {
	fmt.Println("Animal power!", a)
}

func main() {

	p1 := pet{
		"Anonymous",
		true,
		2,
	}
	fmt.Println(p1)

	c1 := cat{
		pet: pet{
			"Bandit",
			true,
			4,
		},
		short_hair: true,
	}
	fmt.Println(c1)
	c1.speak()
	// p1.speak() // can't

	b1 := bird{
		pet: pet{
			"Polly",
			false,
			2,
		},
		tropical: true,
	}
	fmt.Println(b1)
	b1.speak()

	b1.fly()
	// c1.move() // can't
	// p1.move() // can't

	powerCall(b1)
	// powerCall(c1)
	// powerCall(p1) // can't

}
```

### Empty Interface
> every value can be put into an empty interface 

## Assertion 
```go
// Any type that has the method speak on it, is also type animal
type animal interface {
	speak()
	// fly()
}

// both birds and cats are animals because they can both speak
func powerCall(a animal) {
	switch a.(type) {
		case cat:
			fmt.Println("MEOW!!!!!! from", a.(cat).name)
		case bird:
			fmt.Println("CACAW!!!!!! from", a.(bird).name)
	}
	
}
```
