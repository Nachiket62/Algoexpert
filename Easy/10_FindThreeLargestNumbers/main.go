package main

import "fmt"

func main() {

	expected := []int{-2, -1, 7}
	output := FindThreeLargestNumbers([]int{-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7})
	fmt.Println(output)
	fmt.Println(expected)
}