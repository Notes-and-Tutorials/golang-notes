package main
import "fmt"

var y = 42
// var z = "words"

func main() {
	fmt.Print(y) // no line

	fmt.Println(y) // add a line at the end

	fmt.Printf("%T\n", y) // print with formating, so NEEDS a formating string
	fmt.Printf("%b\n", y) // \n escape character for new line
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	fmt.Printf("%#x\t%b\t%x\n", y, y, y) // order matters (note: \t escape character for tab)
	
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // S in front is string print
	fmt.Println(s)
}

// Also F for file print

// https://godoc.org/fmt
/* 
		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value
*/

