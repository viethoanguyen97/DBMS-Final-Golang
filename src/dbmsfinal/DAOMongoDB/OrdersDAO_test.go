package queryMongo

import (
	"fmt"
	"testing"
)

func TestGetOrderInfo(t *testing.T) {
	order, duration, _ := (*OrdersDAO)(nil).GetOrderInfo(150)
	fmt.Println(order)
	fmt.Println(duration)
}

func TestGetAllOrders(t *testing.T) {
	//_, duration, _ := (*OrdersDAO)(nil).GetAllOrdersInfo()
	//fmt.Println(orders)
	//fmt.Println(duration)
}
