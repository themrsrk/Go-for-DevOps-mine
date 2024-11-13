package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
	Age    uint16
}

func UpdateSalary(emp *Employee, newSalary float64) {
	fmt.Println(emp)
	emp.Salary = newSalary
}

func main() {

	emp := Employee{
		Name:   "Shahrukh",
		Salary: 90000,
		Age:    24,
	}
	fmt.Println("Before Updating the vaule = ", emp)

	UpdateSalary(&emp, 70000)

	fmt.Println("After Updating the vaule = ", emp)

}
