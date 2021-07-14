package fibonacci

import (
	"fmt"
)

var f1 int32
var f2 int32

// funс Fibo makes Slice with k first Fibonacci numbers, than prints the sequence.
func FiboSequence(k int) []int32 {
	fmt.Printf("%s %d\n", "Your input is: ", k)
	if k < 0 {
		fmt.Printf("The quantity can not be negative. Please input positive number!\n")
		k = -1
	}

	switch k {
	case -1:
		return nil
	case 0:
		fmt.Printf(" The quantity of Fibonacci numbers = 0!\n")
		return nil
	case 1:
		return []int32{0}
	case 2:
		return []int32{0, 1}
	}
	fiboSlice := make([]int32, 0, k)
	if k > 2 {
		//			fmt.Printf("%s %d \n", "Your input is: ", k)
		//			fmt.Println("0 and 1 are the first and second members of the Fibonacci sequence. Please input number bigger than 2!")

		f1, f2 = 0, 1
		fiboSlice = append(fiboSlice, f1, f2)
		for i := 2; i < k; i++ {
			f1, f2 = f2, f1+f2
			fiboSlice = append(fiboSlice, f2)
		}

	}
	return fiboSlice
}

// funс FiboR calculate k-th number of Fibonacci sequence.
func FiboNumber(k int) (F int) {
	if k < 0 {
		fmt.Printf("A nummer can not be negative. Please input positive number!\n")
	} else {

		if k == 1 {
			F = 0
			return
		} else {
			if k == 2 {
				F = 1

			} else {
				F = FiboNumber(k-1) + FiboNumber(k-2)

			}
		}
	}
	return
}

func FiboSequencePrint(k int) {
	for n, f := range FiboSequence(k) {
		fmt.Println(n+1, ": ", f)
	}
}

func FiboNumberPrint(k int) {
	if k > 0 {
		fmt.Println(k, "-th number of Fibonacci Sequence = ", FiboNumber(k))
	} else {
		fmt.Println("Please input correct number!")
	}
}
