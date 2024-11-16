package main

import (
	"fmt"
)

func modifyArray(arr [5]int, index int, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {

	arrayy := [5]int{2, 4, 5, 1, 3}
	fmt.Println("Original Array: ", arrayy)

	modification := modifyArray(arrayy, 2, 6)
	fmt.Println("Modified Array: ", modification)

	// Let's check if content & length of both the arrays are same or not

	isEqual := arrayy == modification
	fmt.Println("Modified Array: ", isEqual)

	isCopiedEqual := len(arrayy) == len(modification)
	fmt.Println("Modified Array: ", isCopiedEqual)
}
