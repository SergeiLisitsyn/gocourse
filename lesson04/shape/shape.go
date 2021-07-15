package shape

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
