package main

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{
		Height: 5,
		Width:  5,
	}
	DescribeShape(c)
	DescribeShape(r)
}
