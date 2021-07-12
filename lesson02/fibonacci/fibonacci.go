package fibonacci

import (
	"fmt"
)

var f1 int32
var f2 int32

// funс Fibo makes Slice with k first Fibonacci numbers, than prints the sequence.
func Fibo(k int) {
	if k < 0 {
		fmt.Printf("%s %d %s", "Your input is: ", k, " The quantity can not be negative. Please input positive number!")
	} else {
		if k < 3 {
			fmt.Printf("%s %d \n", "Your input is: ", k)
			fmt.Println("0 and 1 are the first and second members of the Fibonacci sequence. Please input number bigger than 2!")
		} else {
			fiboSlice := make([]int32, 0, k)
			f1, f2 = 0, 1
			fiboSlice = append(fiboSlice, f1, f2)
			for i := 2; i < k; i++ {
				f1, f2 = f2, f1+f2
				fiboSlice = append(fiboSlice, f2)
			}
			fmt.Println(fiboSlice)
		}
	}
}

// funс FiboR calculate k-th member of Fibonacci sequence. There is used the recursion method.
func FiboR(k int) (F int) {
	if k == 1 {
		F = 0
		return
	} else {
		if k == 2 {
			F = 1
			return
		} else {
			F = FiboR(k-1) + FiboR(k-2)
			return
		}
	}
}
