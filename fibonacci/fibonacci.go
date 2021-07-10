package fibonacci

import (
	"fmt"
)

var F1 int32
var F2 int32

// funkFibo prints k first Fibonacci numbers
func Fibo(k int) (Fibo []int32) {
	F1, F2 = 0, 1
	Fibo = make([]int32, 1)
	Fibo = append(Fibo, F2)
	//fibo = append(fibo, f1, f2)
	if k > 2 {
		for i := 2; i < k; i++ {
			F1, F2 = F2, F1+F2
			Fibo = append(Fibo, F2)
		}
		return
	} else {
		return
	}
} //Print n memeber of Fibonacci sequence
func ShowFibo(n int) {
	var fiboSlice []int32
	fiboSlice = Fibo(n)
	fmt.Print("The ", n, "-th Fibonacci number = ", fiboSlice[n-1], ".")

} // Print Fibonacci sequence in column
func FiboInColumn(n int) {
	var fiboSlice []int32
	fiboSlice = Fibo(n)
	fmt.Printf("\n%s %s", "Numer", "Fibo number")
	for k, f := range fiboSlice {
		fmt.Printf("\n%5d %5d", k, f)
	}
}
