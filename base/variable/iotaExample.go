package main

import (
	"fmt"
	"reflect"
)

type OrderStatus int

func main() {
	const (
		OrderStatus_Unknown OrderStatus = iota // 0
		OrderStatus_NotActive
		OrderStatus_Untriggered
		OrderStatus_Triggered
		OrderStatus_Active = iota - 1
		OrderStatus_Created
		OrderStatus_Rejected
		OrderStatus_Tested
	)

	fmt.Println(OrderStatus_Unknown, reflect.TypeOf(OrderStatus_NotActive), OrderStatus_NotActive, OrderStatus_Untriggered, OrderStatus_Triggered,
		OrderStatus_Active, OrderStatus_Created, OrderStatus_Rejected, OrderStatus_Tested)
}
