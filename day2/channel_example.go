package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func processedOrders(orderID int, statusChannel chan Order) {
	time.Sleep(time.Second * 2)

	statusChannel <- Order{ID: orderID, Status: "Completed"} //This means we are processing the orders and sending it thorugh the channel
}

func main() {
	statusChannel := make(chan Order)

	for i := 1; i <= 5; i++ {
		go processedOrders(i, statusChannel)
	}

	// Receiving and printing order statuses
	for i := 1; i <= 5; i++ {
		processedOrder := <-statusChannel
		fmt.Printf("Order %d processed: %s\n", processedOrder.ID, processedOrder.Status)
	}
}
