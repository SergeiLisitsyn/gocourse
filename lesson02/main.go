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
	fmt.Println("You entered:", f)
	// f = 10 // Size of Fibonacci sequence
	if f > 0 {
		defer fmt.Print("\nNow you see sequence of ", f, " Fibonacci numbers.")
	}
	fibonacci.Fibo(f)
	fmt.Println()
	for k := 1; k <= f; k++ {
		fmt.Println(k, " ", fibonacci.FiboR(k))
	}

}
