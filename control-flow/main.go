/* control flow (how computer reads code):
1 sequence
2 loop; iterative
3 conditional
*/

// Loops
// NOTE: There is no while loop in Go

package main

import "fmt"

func main() {
	// for init; condition; post {}

	// for i := 0; i <= 10; i++ {
	// 	fmt.Printf("The outer loop: %d\n", i)
	// 	for j:=  0; j < 3; j++ {
	// 		fmt.Printf("\t\tThe inner loop: %d\n", j)
	// 		}
	// }

	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }

	// could use to create an eternal loop to always be listening for something
	// x := 1
	// for {
	// 	if x > 9 {
	// 		break
	// 	}
	// 	fmt.Println(x)
	// 	x++
	// }
	// fmt.Println("done.")

	// x := 0
	// for {
	// 	x++
	// 	if x > 100 {
	// 		break
	// 	}

	// 	if x%2 != 0 {
	// 		continue
	// 	}

	// 	fmt.Println(x)
	// }

	// Print ascii characters
	// for i := 33; i <= 122; i++ {
	// 	fmt.Printf("%v\n\t%#U\n", i, i)
	// }

	//  if true { fmt.Println("001") }
	//  if false { fmt.Println("002") }
	//  if !true { fmt.Println("003") }
	//  if !false { fmt.Println("004") }

	// Go needs semicolons at the end of all statements (usually will do in compile for you based on line and context, but in this case you need it because 2 statements on 1 line here)
	// Initialization statement
	// Limits scope
	// if x := 12; x == 2 {
	// 	fmt.Println("statement true")
	// }
	// fmt.Println("statement false")

	// x := 42
	// if x == 40 {
	// 	fmt.Println("value was 40")
	// } else {
	// 	fmt.Println("value was NOT 40")
	// }

	// x := 42
	// if x == 40 {
	// 	fmt.Println("value was 40")
	// } else if x == 41 {
	// 	fmt.Println("value was 40")
	// } else {
	// 	fmt.Println("value was NOT 40")
	// }

	// Switch statement
	// i := 2
	// fmt.Print("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one") 
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// Fallthrough
	// Default for switch statements are to stop checking when reach a true.
	// Fallthrough makes it continue to run
	// Generally don't use
	// switch {
	// case false:
	// 	fmt.Println("this should not print")
	// case (2 == 4):
	// 	fmt.Println("this should not print2")
	// case (3 == 3):
	// 	fmt.Println("prints")
	// 	fallthrough
	// case (4 == 4):
	// 	fmt.Println("also true, does it print?")
	// 	fallthrough
	// case (7 == 9):
	// 	fmt.Println("not true1")
	// 	fallthrough
	// case (11 == 14):
	// 	fmt.Println("not true 2")
	// 	fallthrough
	// case (15 == 15):
	// 	fmt.Println("true 15")
	// 	fallthrough
	// default:
	// 	fmt.Println("this is default")
	// }

	// n := "Bond"
	// switch n {
	// case "Moneypenny":
	// 	fmt.Println("miss money")
	// case "Bond":
	// 	fmt.Println("bond james")
	// case "Q":
	// 	fmt.Println("this is q")
	// default:
	// 	fmt.Println("this is default")
	// }

	// n := "Bond"
	// switch n {
	// case "Moneypenny", "Bond", "Do No":
	// 	fmt.Println("miss money or bond or dr no")
	// case "M":
	// 	fmt.Println("m")
	// case "Q":
	// 	fmt.Println("this is q")
	// default:
	// 	fmt.Println("this is default")
	// }
}
