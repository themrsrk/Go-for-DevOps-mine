package main

import "fmt"

func main() {
	// Basic type variabls
	// int, float, bool, string etc
	// 1. Signed Integers (int8, int16, int32, int64) depending on the platform
	// 2. Unsigned Integers (uint8 or byte, uint16, uint32, uint64), size is platform dependent
	// 3. Machine Dependent Types (int, uint, uintptr) depends on the type of architecture on which the program is running 32bit vs 64bit

	//Composite type variables
	// struct, array, slices, map, channels, interfaces, pointers etc

	fmt.Println("\nSigned Integers")
	var a int8 = 127     //ranges from -128 to 127
	var b int16 = -32768 //ranges from -32768 to 32767
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println("\nUnsigned Integers ")
	var c uint8 = 255    //ranges from 0 to 255
	var d uint16 = 65535 //ranges from 65535
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println("\nMachine dependent Types")
	var e int = 91839123980
	var f uint = 7831971398
	var g uintptr = 0x7383233
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)

	fmt.Println("\nFloating Types")
	var h float32 = 3.1456
	var i float64 = 3.14577898309283
	fmt.Println("h=", h)
	fmt.Println("i=", i)

	fmt.Println("\nComplex Numbers")
	var j complex64 = 1 + 2i
	var k complex128 = 2 + 2i
	fmt.Println("j=", j)
	fmt.Println("k=", k)

	fmt.Println("\nBool")
	var l bool = true
	var m bool = false
	fmt.Println("l=", l)
	fmt.Println("m=", m)

	fmt.Println("\nStrings") //Strings in go are immutable
	var n string = "Shahrukh Siddiqui"
	fmt.Println("n=", n)
}
