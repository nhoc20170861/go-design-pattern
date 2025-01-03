package main

import "fmt"

// Strategy interface defines the execute method that all strategies must implement
type Strategy interface {
	execute(a, b int) int
}

// Add struct implements the Strategy interface for addition
type Add struct{}

// execute performs addition
func (add *Add) execute(a, b int) int {
	return a + b
}

// Subtract struct implements the Strategy interface for subtraction
type Subtract struct{}

// execute performs subtraction
func (s *Subtract) execute(a, b int) int {
	return a - b
}

// Multiply struct implements the Strategy interface for multiplication
type Multiply struct{}

// execute performs multiplication
func (m *Multiply) execute(a, b int) int {
	return a * b
}

// Calculator struct holds a reference to a Strategy
type Calculator struct {
	strategy Strategy
}

// executeStrategy executes the strategy's execute method
func (c *Calculator) executeStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

// setStrategy sets the strategy for the calculator
func (c *Calculator) setStrategy(st Strategy) {
	c.strategy = st
}

func main() {
	c := &Calculator{}

	add := &Add{}
	subtract := &Subtract{}
	mul := &Multiply{}

	// Set strategy to Add and execute
	c.setStrategy(add)
	ret := c.executeStrategy(3, 4)
	fmt.Printf("Addition result: %v\n", ret)

	// Set strategy to Subtract and execute
	c.setStrategy(subtract)
	ret = c.executeStrategy(5, 1)
	fmt.Printf("Subtraction result: %v\n", ret)

	// Set strategy to Multiply and execute
	c.setStrategy(mul)
	ret = c.executeStrategy(4, 3)
	fmt.Printf("Multiplication result: %v\n", ret)
}
