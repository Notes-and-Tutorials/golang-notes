```go
type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (a square) area() float64 {
	return a.side * a.side

}

func (b circle) area() float64 {
	return b.radius * b.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(c shape) float64{
	
	return c.area()
}

func main() {
	s1 := square{3}
	c1 := circle{2}

	fmt.Println(s1.area())
	fmt.Println(c1.area())
	
	x := info(s1)
	fmt.Println(x)
	
	y := info(c1)
	fmt.Println(y)
	
}
```
