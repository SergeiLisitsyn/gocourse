package main

import (
	"fmt"

	"github.com/SergeiLisitsyn/gocourse/fibonacci"
)

func main() {
	var f int = 12 // Quontity of Fibonacci numbers
	var fiboArr []int32
	fiboArr = fibonacci.Fibo(f)
	fmt.Println(fiboArr)

}
