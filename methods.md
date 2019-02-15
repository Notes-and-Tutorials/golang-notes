 # Method

 ```go
 type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

// func (r receiver) indentifier(parameters)(returns(s)){code block}

// method
// any value of that type (secretAgent) has access this method (speak)
func (s secretAgent) speak() {
	fmt.Println("i am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
			32,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak() // calling the method

} // end main
```