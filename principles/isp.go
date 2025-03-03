package principles

import "fmt"

// ISP - Interface Segregation Principle
// - A client should never be forced to implement an interface that it doesn't use
// - It is better to have many specific interfaces than one general interface
// - It makes the code more readable and maintainable
// Avoid large interfaces that force unused methods.
// Use smaller, focused interfaces that match specific behaviors.
// Improve maintainability by reducing unnecessary dependencies.

type TwoDShape interface {
	Area() float64
}

type ThreeDShape interface {
	TwoDShape
	Volume() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Cube struct {
	Length, Breadth, Height float64
}

func (c Cube) Area() float64 {
	return 6 * c.Length * c.Breadth
}

func (c Cube) Volume() float64 {
	return c.Length * c.Breadth * c.Height
}

func ISP() {
	fmt.Println("From ISP Package!!!")
	square := Square{Side: 5}
	cube := Cube{Length: 5, Breadth: 5, Height: 5}
	fmt.Printf("Area of Square: %.2f\n", square.Area())
	fmt.Printf("Area of Cube: %.2f\n", cube.Area())
	fmt.Printf("Volume of Cube: %.2f\n", cube.Volume())
}
