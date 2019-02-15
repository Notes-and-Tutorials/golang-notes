package main

import "fmt"

/* control flow (how computer reads code):
1 sequence
2 loop; iterative
3 conditional
*/

func main() {
	fmt.Println("hello anything") // 1

	foo() // 2 (see below)

	fmt.Println("after foo") // 3

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i) // 4
		}
	}

	bar() //5 (see below)
}

func foo() {
	fmt.Println("in foo") // 2
}

func bar() {
	fmt.Println("exited") // 5
}
