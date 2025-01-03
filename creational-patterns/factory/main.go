package main

import "fmt"

// Shape is an interface with a Draw method
type Shape interface {
	Draw()
}

// Circle is a struct that implements the Shape interface
type Circle struct{}

// Draw prints a message indicating that a Circle is drawing
func (c *Circle) Draw() {
	fmt.Println("Circle is drawing")
}

// Rectangle is a struct that implements the Shape interface
type Rectangle struct{}

// Draw prints a message indicating that a Rectangle is drawing
func (r *Rectangle) Draw() {
	fmt.Println("Rectangle is drawing")
}

// CreateShape returns a Shape based on the provided type
func CreateShape(t string) Shape {
	switch t {
	case "circle":
		return &Circle{}
	case "rectangle":
		return &Rectangle{}
	default:
		fmt.Println("This shape is not supported")
		return nil
	}
}

func main() {
	shape1 := CreateShape("circle")
	if shape1 != nil {
		shape1.Draw()
	}

	shape2 := CreateShape("rectangle")
	if shape2 != nil {
		shape2.Draw()
	}

	shape3 := CreateShape("triangle")
	if shape3 != nil {
		shape3.Draw()
	}
}
