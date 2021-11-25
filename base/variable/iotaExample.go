package main

import "fmt"

type OrderStatus int

func main() {
	const (
		OrderStatus_Unknown OrderStatus = iota //0
		OrderStatus_NotActive
		OrderStatus_Untriggered
		OrderStatus_Triggered = 3
		OrderStatus_Active    = 3
		OrderStatus_Created   = 4
		OrderStatus_Rejected  = iota
		OrderStatus_Tested
	)
	fmt.Println(OrderStatus_Unknown, OrderStatus_NotActive, OrderStatus_Untriggered, OrderStatus_Triggered,
		OrderStatus_Active, OrderStatus_Created, OrderStatus_Rejected, OrderStatus_Tested)
}
