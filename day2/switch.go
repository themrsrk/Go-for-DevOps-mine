package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Basic switch:")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	// ...
	default:
		fmt.Println("Invalid day")
	}

	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday: //we can use multiple conditions either this or this
		fmt.Println("It's Weekend")
	default:
		fmt.Println("Its a Weekday")
	}

	a := 10
	switch {
	case a > 11:
		fmt.Println("Greater than 9")
	case a == 10:
		fmt.Println("Equal to 10")
	default:
		fmt.Println("Nothing")
	}

}
