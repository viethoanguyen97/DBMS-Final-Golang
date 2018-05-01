package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type Orderdetail struct {
	OrderID       int `json:"order_id"`
	CarID         int `json:"car_id"`
	QuantityOrder int `json:"quantity_order"`
}

type ODetail struct {
	CarID         int `json:"car_id"`
	QuantityOrder int `json:"quantity_order"`
}

type OCar struct {
	OrderID int `json:"order_id"`
	CarID   int `json:"car_id"`
}

type OrderWithDetails struct {
	OrderID int        `json:"order_id"`
	Details []*ODetail `json:"details"`
}

type Orderdetails []*Orderdetail

func genOrderdetail(order_id int) ([]*Orderdetail, *OrderWithDetails) {
	numberOfOrderdetails := rand.Intn(100) + 1

	orderdetails := make([]*Orderdetail, 0)

	details := make([]*ODetail, 0)

	m := make(map[OCar]bool)
	for i := 0; i < numberOfOrderdetails; i++ {
		ocar := OCar{}
		detail := &ODetail{}
		order_detail := &Orderdetail{}

		for {
			carID := rand.Intn(1000) + 1
			quantityOrder := rand.Intn(100) + 1

			ocar = OCar{
				CarID:   carID,
				OrderID: order_id,
			}

			if m[ocar] == false {
				m[ocar] = true
				detail = &ODetail{
					CarID:         carID,
					QuantityOrder: quantityOrder,
				}

				order_detail = &Orderdetail{
					OrderID:       order_id,
					CarID:         carID,
					QuantityOrder: quantityOrder,
				}

				break
			}
		}

		orderdetails = append(orderdetails, order_detail)
		details = append(details, detail)
	}

	orderwithdetails := &OrderWithDetails{
		OrderID: order_id,
		Details: details,
	}

	return orderdetails, orderwithdetails
}

func main() {
	orders := Orderdetails{}
	orders_with_details := make([]*OrderWithDetails, 0)

	for i := 1; i <= 1000; i++ {
		i_orders_details, i_orders_with_details := genOrderdetail(i)
		orders = append(orders, i_orders_details...)
		orders_with_details = append(orders_with_details, i_orders_with_details)
	}

	//write to json file
	orders_json, _ := json.Marshal(orders)
	err := ioutil.WriteFile("orderdetails.json", orders_json, 0644)
	if err != nil {
		panic(err)
	}

	orders_with_details_json, _ := json.Marshal(orders_with_details)
	err = ioutil.WriteFile("orderwithdetails.json", orders_with_details_json, 0644)
	if err != nil {
		panic(err)
	}
}

