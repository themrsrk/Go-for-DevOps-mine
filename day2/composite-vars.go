package main

import "fmt"

func main() {

	// 1. Array (An array has a fixed size defined at compile time e:g var a[5] int will declare an array of 5 integers)
	fmt.Println("Array")
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("array = ", a)

	// 2. Slices (Slices are similar to array but their size can change and they are more flexible)
	fmt.Println("\nSlices")
	b := []int{2, 4, 5, 2, 4, 9}
	c := []string{"shahrukh", "siddiqui"}
	fmt.Println("slices int = ", b)
	fmt.Println("slices string = ", c)

	// 3. Structs (Structs are used to create custom data types with mixed data, for example type Person struct {Name string, Age int})
	fmt.Println("\nStructs")
	type Person struct {
		Name string
		Age  int
	}

	var q Person = Person{"Shahrukh", 24}
	fmt.Println("Struct ", q)
	fmt.Println("Struct Name ", q.Name)
	fmt.Println("Struct Age", q.Age)

	// 4. Maps (Maps are used to store key value pairs, same as dictionary in python e:g a : = map [key]value {key : value})
	fmt.Println("\nMaps")
	r := map[int]string{1: "Shahrukh", 2: "Siddiqui"}
	fmt.Println("Maps ", r)
	fmt.Println("Maps ", r[1])
	fmt.Println("Maps ", r[2])

	// 5. Pointers (Holds the memory address of the variable for example p *int is a pointer to an integer)
	fmt.Println("\nPointers")
	s := 10
	t := &s               //& sign is used to get the memory address of integer
	fmt.Println("S", s)   //holds the original variable
	fmt.Println("T", t)   //holds the memory address where s is stored
	fmt.Println("*T", *t) //print the value of variable at the memory address stored in t (basically it's fetching the value from address)

	// 6. Channels (Channels are used for communication between goroutines (concurrent threads), they can be synchronouse and asynchronous)
	fmt.Println("\nChannels")

	//basic example of channel
	u := make(chan int) //unbuffered channel of integers
	fmt.Println("Channel ", u)

}
