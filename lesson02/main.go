package main

import (
	"fmt"

	"github.com/SergeiLisitsyn/gocourse/lesson02/fibonacci"
)

func main() {
	defer fmt.Println("\nThe Progam has successfully completed!") //Complete massage
	fmt.Println("Hello, please input f - qauntity of Fibonacci numbers you want to print")
	var f int
	fmt.Scanf("%d", &f)
	//f = 14 // Size of Fibonacci sequence

	fibonacci.FiboSequencePrint(f)
	fibonacci.FiboNumberPrint(f)

}
