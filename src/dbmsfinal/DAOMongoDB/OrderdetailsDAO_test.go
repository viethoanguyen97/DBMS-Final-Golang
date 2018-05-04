package queryMongo

import (
	"testing"
)

func TestGetOrderdetailsOfOrderID(t *testing.T) {
	//_, duration, _ := (*OrderDetailsDAO)(nil).GetOrderdetailsOfOrderID(1)
	//fmt.Println(duration)
}

func TestGetOrderCardetailsOfOrderID(t *testing.T) {
	//_, duration, _ := (*OrderDetailsDAO)(nil).GetOrderCardetailsOfOrderID(1)
	//fmt.Println(duration)
}

func TestGetOrderCardetailsOfCustomerID(t *testing.T) {
	/*_, duration, _ := (*OrderDetailsDAO)(nil).GetOrderCardetailsOfCustomerID(135)
	fmt.Println(duration)*/
}

func TestGetInsertOrderdetails(t *testing.T) {
	/*insertNewOrderdetailData := &dataMongoDB.OrderdetailCSV{
		OrderID:       1,
		CarID:         1001,
		QuantityOrder: 1,
	}

	_, duration, err := (*OrderDetailsDAO)(nil).InsertNewOrderdetail(insertNewOrderdetailData)

	fmt.Println(err)
	fmt.Println(duration)*/
}

func TestGetDeleteOrderdetails(t *testing.T) {
	/*duration, err := (*OrderDetailsDAO)(nil).DeleteOrderdetail(1, 1001)

	fmt.Println(err)
	fmt.Println(duration)*/
}
