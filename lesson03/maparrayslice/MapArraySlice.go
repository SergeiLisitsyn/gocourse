package maparrayslice

import (
	"fmt"
	"sort"
)

// func sums values all members of array and divide on length of array
func ArrayAverage(arr []float32) (arrAvarage float32) {
	var sum float32 = 0
	for _, v := range arr {
		sum = sum + v
	}
	return sum / float32(len(arr))
}

// Func finds longest string in strind array
func LongestString(sliceStr []string) {
	max := 0                     // defoult max length
	mStr := ""                   // defoult string with max length
	for _, v := range sliceStr { //for loop to find string with max length
		if len(v) > max {
			max = len(v)
			mStr = v
		}
	}
	fmt.Println(mStr)
} // Func returns the copy of the original slice in reverse order.
func Reverse(testSlice []int64) (reverseSlice []int64) {
	reverseSlise := make([]int64, 0, len(testSlice)) // new slice foe elements in reverse oder.
	for k := len(testSlice); k > 0; k-- {
		reverseSlise = append(reverseSlise, testSlice[k-1]) // loop for adding elements
	}
	return reverseSlise
}

func MapSortByKey(testMap map[int]string) {
	var sliceMap = make([]int, 0, len(testMap)) // Slice for map keys.
	for k := range testMap {                    // Loop for adding keys as elements
		sliceMap = append(sliceMap, k)
	}
	sort.Ints(sliceMap) // There is being used Func sort.Ints() for sorting int slice with map keys.
	for _, v := range sliceMap {
		fmt.Println(v, testMap[v]) // Range loop prints sorted keys and map values.
	}
}
