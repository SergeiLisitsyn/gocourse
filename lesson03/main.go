package main

import (
	"fmt"

	"github.com/SergeiLisitsyn/gocourse/lesson03/maparrayslice"
)

func main() {
	defer fmt.Println("All tasks have been completed.")
	fmt.Println("Hello! There is the HoneWork for lesson03")
	fmt.Println()

	fmt.Println("Task1")
	testArr := [...]float32{1, 2, 3, 4, 5, 6}
	fmt.Println(maparrayslice.ArrayAverage(testArr[:]))
	fmt.Println()

	fmt.Println("Task2")
	stringSlice := []string{"one", "two", "three"}
	maparrayslice.LongestString(stringSlice)
	fmt.Println()
	stringSlice = stringSlice[:2]
	maparrayslice.LongestString(stringSlice)
	fmt.Println()

	fmt.Println("Task3")
	testSlice := []int64{1, 2, 5, 15}
	fmt.Println(testSlice)
	reversSlice := maparrayslice.Reverse(testSlice)
	fmt.Println(reversSlice)
	fmt.Println()

	fmt.Println("Task4")
	var testMap1 = map[int]string{2: "a", 0: "b", 1: "c"}
	var testMap2 = map[int]string{10: "aa", 0: "bb", 500: "cc"}
	maparrayslice.MapSortByKey(testMap1)
	fmt.Println()
	maparrayslice.MapSortByKey(testMap2)

}
