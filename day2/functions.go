//Block of Code that can use multiple times

package main

import "fmt"

func add(a, b int) int { //here we are defining that bot a and b will be type of int and return type is also int
	c := a + b
	return c
}

func swap(c, d string) (string, string) { //here we are showing that we can use multiple return datatypes
	return d, c
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	addition := add(5, 5)
	fmt.Println(addition)

	swapping1, swapping2 := swap("Shahrukh", "Siddiqui")
	fmt.Println(swapping1, swapping2)

	Sum := sum(2, 3, 4, 5, 6)
	fmt.Println(Sum)
}
