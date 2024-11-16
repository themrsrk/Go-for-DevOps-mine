package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Hello Func Ended")
}

func sayHi() {
	fmt.Println("Hi World")
}

func main() {
	fmt.Println("Hi Main")
	go sayHello() //Whenever we put go keyword infront of any function it becomes asynchornous means now it will be executed concurrently not sequentially
	//Like Main prorgam won't be waiting for this function

	sayHi()

	time.Sleep(4 * time.Second)
}
