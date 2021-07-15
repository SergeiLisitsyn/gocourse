package main

import (
	"github.com/SergeiLisitsyn/gocourse/lesson04/circle"
	"github.com/SergeiLisitsyn/gocourse/lesson04/rectangle"
	"github.com/SergeiLisitsyn/gocourse/lesson04/shape"
)

func main() {

	c := circle.Circle{Radius: 10}
	r := rectangle.Rectangle{
		Height: 8,
		Width:  6,
	}

	shape.DescribeShape(c)
	shape.DescribeShape(r)
}
