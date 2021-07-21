package main

import (
	"fmt"
	"testing"
)

func TestDescribeShape(t *testing.T) {
	r0 := Rectangle{
		Height: 0,
		Width:  0,
	}
	r5 := Rectangle{
		Height: 5,
		Width:  5,
	}
	c0 := Circle{Radius: 0}
	c1 := Circle{Radius: 1}

	negativeCircleResult := DescribeShape(c0)
	positiveCircleResult := DescribeShape(c1)
	negativeRectangleResult := DescribeShape(r0)
	positiveRectangleResult := DescribeShape(r5)
	fmt.Printf("\nHello, %v", DescribeShape(c0))
	// test for side length <= 0
	if negativeRectangleResult != "\nError. Rectangle with sides = 0 , 0 doesn't exist." {
		t.Errorf("String() failed, expected %v, got %v", "Error.", negativeRectangleResult)
	}
	// test for sides = 5
	if positiveRectangleResult != "\nRectangle with height 5.00 and width 5.00\nArea: 25.00\nPerimeter: 20.00" {
		t.Errorf("String() failed, expected %v, got %v", "\nRectangle with height 5.00 and width 5.00\nArea: 25.00\nPerimeter: 20.00\n", positiveRectangleResult)
	}

	// test for radius<= 0
	if negativeCircleResult != "\nError. Circle with radius = 0.00 doesn't exist." {
		t.Errorf("String() failed, expected %v, got %v", "Error. Circle with Radius = 0 doesn't exist.", negativeCircleResult)
	}
	// test for radius = 1
	if positiveCircleResult != "\nCircle: radius 1.00\nArea: 3.14\nPerimeter: 6.28" {
		t.Errorf("String() failed, expected %v, got %v", "\nCircle: radius 1.00", positiveCircleResult)
	}
}
