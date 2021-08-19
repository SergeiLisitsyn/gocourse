package main

import (
	"errors"
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) String() string {
	if c.Radius <= 0 {
		return fmt.Sprintf("\nError! Circle with Radius = %v doesn't exist.", c.Radius)
	}
	return fmt.Sprintf("\nCircle: radius %.2f", c.Radius)
}

func (c Circle) Area() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("Error is happen! Radius could not be <= 0.")
	}
	return math.Pi * math.Pow(c.Radius, 2), nil
}
func (c Circle) Perimeter() (float64, error) {
	if c.Radius <= 0 {
		return 0, errors.New("Error is happen! Radius could not be <= 0.")
	}
	return 2 * math.Pi * c.Radius, nil
}
