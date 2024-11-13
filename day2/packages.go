//Package in go are a way to organize and encapsulate the code, Each package can contain multiple go files and each file belongs to one package
//We just need to write at the beggining oft he code

//Exported names (In golang a name is exported if it begins with a Capital letter for example, Printer is an exported name and can be used from other packages, while printer is not)
//You can use code from other packages by importing them

package main

import (
	"day2/shapes"
	"fmt"
)

func main() {

	radius := 5.0
	radius_result := shapes.AreaOfCircle(radius)
	fmt.Println("Area Of circle is: ", radius_result)

	side := 9.0
	sq_result := shapes.AreaOfSquare(side)
	fmt.Println("Area Of Square is: ", sq_result)
}
