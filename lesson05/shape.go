package main

import (
	"fmt"
)

type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	a, _ := s.Area()
	fmt.Printf("Area: %.2f \n", a)
	p, _ := s.Perimeter()
	fmt.Printf("Perimeter: %.2f\n", p)
}
