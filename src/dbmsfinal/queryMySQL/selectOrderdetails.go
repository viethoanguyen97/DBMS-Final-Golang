package queryMySQL

import (
	"errors"
	"fmt"
	"time"
)

func (r *OrderdetailsDAO) GetOrderdetailsOfOrderID(order_id int64) ([]*Orderdetail, float64, error) {
	orderdetails := make([]*Orderdetail, 0)

	start := time.Now()

	rows, err := db.Query("SELECT o.order_id, o.customer_id, od.car_id, od.quantity_order FROM Orders o "+
		"JOIN Orderdetails od ON o.order_id = od.order_id "+
		"WHERE o.order_id = ?;", order_id)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetail of order_id")
	}

	cnt := 0
	for rows.Next() {
		orderdetail := &Orderdetail{}
		cnt++
		err := rows.Scan(&orderdetail.OrderID, &orderdetail.CustomerID, &orderdetail.CarID, &orderdetail.QuantityOrder)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get orderdetails of orderid")
		}

		orderdetails = append(orderdetails, orderdetail)
	}

	defer rows.Close()

	return orderdetails, elapsed, nil
}

func (r *OrderdetailsDAO) GetOrderCardetailsOfOrderID(order_id int64) ([]*OrderCarDetail, float64, error) {
	ordercardetails := make([]*OrderCarDetail, 0)

	start := time.Now()

	rows, err := db.Query("SELECT o.order_id, o.customer_id, od.car_id, c.car_model, od.quantity_order FROM Orders o "+
		"JOIN Orderdetails od ON o.order_id = od.order_id "+
		"JOIN Cars c ON od.car_id = c.car_id "+
		"WHERE o.order_id = ?;", order_id)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetail of order_id")
	}

	cnt := 0
	for rows.Next() {
		ordercardetail := &OrderCarDetail{}
		cnt++
		err := rows.Scan(&ordercardetail.OrderID, &ordercardetail.CustomerID, &ordercardetail.CarID, &ordercardetail.CarModel, &ordercardetail.QuantityOrder)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get orderdetails of orderid")
		}

		ordercardetails = append(ordercardetails, ordercardetail)
	}

	defer rows.Close()

	return ordercardetails, elapsed, nil
}

func (r *OrderdetailsDAO) GetOrderCardetailsOfCustomerID(customer_id int64) ([]*OrderCarDetail, float64, error) {
	ordercardetails := make([]*OrderCarDetail, 0)

	start := time.Now()

	rows, err := db.Query("SELECT o.order_id, o.customer_id, od.car_id, c.car_model, od.quantity_order FROM Orders o "+
		"JOIN Orderdetails od ON o.order_id = od.order_id "+
		"JOIN Cars c ON od.car_id = c.car_id "+
		"WHERE o.customer_id = ?;", customer_id)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetail of order_id")
	}

	cnt := 0
	for rows.Next() {
		ordercardetail := &OrderCarDetail{}
		cnt++
		err := rows.Scan(&ordercardetail.OrderID, &ordercardetail.CustomerID, &ordercardetail.CarID, &ordercardetail.CarModel, &ordercardetail.QuantityOrder)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get orderdetails of orderid")
		}

		ordercardetails = append(ordercardetails, ordercardetail)
	}

	defer rows.Close()

	return ordercardetails, elapsed, nil
}
