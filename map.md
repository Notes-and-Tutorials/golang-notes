# Map
- key value store
- great for fast look up
- unordered list

## The comma ok idiom

```go
func main() {
	// composite literal
	// [KEY type]VALUE
	// `map[string]int` is the entire type
	m := map[string]int{
		"Hailee": 27,
		"Nikko":  26, // need trailing comma
	}
	fmt.Println(m) // prints:map[Hailee:27 Nikko:26]

	// access
	fmt.Println(m["Nikko"])
	fmt.Println(m["no one"]) // prints: 0 // whatever the 0 value is for that type

	// check for existance in map
	// the comma ok idiom
	v, ok := m["no one"]
	fmt.Println(v)  // print value
	fmt.Println(ok) // print whether ok

	m["Bandit"] = 1 // * add new element

	// conditional
	if v, ok := m["no one"]; ok {
		fmt.Println("Age is ", v)
	} else {
		fmt.Println("Not in database")
	}

	// * range over
	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
```

## Map add element & range

```go
func main() {
	// composite literal
	// [KEY type]VALUE
	// `map[string]int` is the entire type
	m := map[string]int{
		"Hailee": 27,
		"Nikko":  26, // need trailing comma
	}
	fmt.Println(m) // prints:map[Hailee:27 Nikko:26]

	// access
	fmt.Println(m["Nikko"])
	fmt.Println(m["no one"]) // prints: 0 // whatever the 0 value is for that type

	// check for existance in map
	// the comma ok idiom
	v, ok := m["no one"]
	fmt.Println(v)  // print value
	fmt.Println(ok) // print whether ok

	m["Bandit"] = 1 // * add new element

	// conditional
	if v, ok := m["no one"]; ok {
		fmt.Println("Age is ", v)
	} else {
		fmt.Println("Not in database")
	}

	// * range over
	for k, v := range m {
		fmt.Println(k, v)
	}
}

```

## Delete an entire from a map
` delete(<map name>, "key")`
```go
func main() {
	m := map[string]int{
		"Hailee": 27,
		"Nikko":  26,
	}
	fmt.Println(m)

	delete(m, "Hailee")
	fmt.Println(m)

	// Does not throw error if does not exist
	delete(m, "Bandit")
	fmt.Println(m)

	// So use comma ok idiom
	// Note `m["Bandit"]` is called a lookup
	if v, ok := m["Bandit"]; ok {
		fmt.Println("value:", v)
		delete(m, "Bandit")
	} else {
		fmt.Println("does not exist")
	}

}
```
#### map with a value type slice
```go
func main() {
	b := map[string][]string{
		"Miu_Hailee": []string{"bread", "toast"},
		"Miu_Nikko": []string{"nothing", "everything"},
	}

	fmt.Println(b)

}
```
