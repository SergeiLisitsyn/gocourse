package main

import (
	"fmt"

	"github.com/SergeiLisitsyn/gocourse/fibonacci"
)

func main() {
	defer fmt.Println("\nThe Progam has successfully completed!") //Complete massage
	fmt.Println("Hello, please input f - qauntity of Fibonacci numbers you want to print")
	var f int
	fmt.Scanf("%d", &f)
	fmt.Println("You entered:", f)
	//	f = 12 // Size of Fibonacci sequence
	defer fmt.Print("\nNow you see sequence of ", f, " Fibonacci numbers.")
	var fiboArr []int32
	fiboArr = fibonacci.Fibo(f)
	fmt.Println(fiboArr)
	fmt.Print()
	fibonacci.ShowFibo(10)
	fibonacci.FiboInColumn(f)

}
