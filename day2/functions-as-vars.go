package main

import "fmt"

func main() {

	var a func(int) int
	a = func(x int) int { return x * x }
	b := a(8)
	result := b //result will be 64
	fmt.Println("Result ", result)
}
