// Go has only one loop construct which is for loop, you can use it in different way
package main

import "fmt"

func main() {

	// fmt.Println("Basic for loop")
	// for a := 1; a <= 20; a += 2 {
	// 	fmt.Println(a)
	// }

	// fmt.Println("For loop as a while")
	// a := 5
	// for a <= 20 {
	// 	fmt.Println(a)
	// 	a++
	// }

	// fmt.Println("Infinite Loop")
	// a := 200
	// for {
	// 	if a < 100 {
	// 		break
	// 	}
	// 	fmt.Println(a)
	// 	a--
	// }

	// fmt.Println("Loops on array or slies")

	// a := []int{1, 2, 3, 4, 5, 6, 7}

	// for key, val := range a {
	// 	fmt.Printf("Index: %d, Value: %d\n", key, val)
	// }

	fmt.Println("Loops on array or Strings")
	b := "ShahrukhSiddiqui"
	for key, val := range b {
		fmt.Printf("Index: %d, Value: %c\n", key, val)
	}

}
