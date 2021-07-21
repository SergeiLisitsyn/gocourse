package main

import (
	"fmt"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

func DescribeShape(s Shape) string {
	a, _ := s.Area()
	p, err := s.Perimeter()
	Str := "kkk"
	//str := fmt.Sprintf("%v", s)
	if err == nil {
		fmt.Printf("\n%v\nArea: %.2f\nPerimeter: %.2f\n", s, a, p)
		return fmt.Sprintf("%s\nArea: %.2f\nPerimeter: %.2f", s, a, p)
	} else {

		switch v := s.(type) {
		case Circle:
			Str = fmt.Sprintf("\nError. Circle with radius = %.2f doesn't exist.", v.Radius)
			fmt.Printf(Str)
		case Rectangle:
			Str = fmt.Sprintf("\nError. Rectangle with sides = %v , %v doesn't exist.", v.Height, v.Width)
			fmt.Printf(Str)
		}
		return Str
	}
}
