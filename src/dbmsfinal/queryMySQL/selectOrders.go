package queryMySQL

import (
	"errors"
	"fmt"
	"time"
)

func (r *OrderDAO) GetOrderInfo(order_id int64) (*Order, float64, error) {
	orderInfo := &Order{}

	start := time.Now()

	row := db.QueryRow("SELECT order_id, customer_id FROM Orders WHERE order_id = ?;", order_id)

	elapsed := time.Since(start).Seconds()

	err := row.Scan(&orderInfo.OrderID, &orderInfo.CustomerID)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get order info")
	}

	return orderInfo, elapsed, nil
}

func (r *OrderDAO) GetAllOrdersInfo() ([]*Order, float64, error) {
	orders := make([]*Order, 0)

	start := time.Now()

	rows, err := db.Query("SELECT order_id, customer_id FROM Orders")

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orders info")
	}

	cnt := 0
	for rows.Next() {
		orderInfo := &Order{}
		cnt++
		err := rows.Scan(&orderInfo.OrderID, &orderInfo.CustomerID)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get all orders info")
		}

		orders = append(orders, orderInfo)
	}

	defer rows.Close()

	return orders, elapsed, nil
}
