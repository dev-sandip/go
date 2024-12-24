package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func newOrder(id string, amount float32, status string) *order {
	order := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &order
}

// user defined struct methods
func (o *order) changeStatus(status string) {
	o.status = status
}
func main() {
	// myOrder := order{
	// 	id: "1", amount: 50, status: "Received",
	// }
	// myOrder2 := order{
	// 	id: "2", amount: 55.00, status: "Received", createdAt: time.Now(),
	// }
	// myOrder2.changeStatus("Paid")
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder)
	// fmt.Println(myOrder2)
	myOrder := newOrder("1", 55.21, "Paid")
	fmt.Println(myOrder)
}
