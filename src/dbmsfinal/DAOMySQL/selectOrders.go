package queryMySQL

import (
	"errors"
	"fmt"
	"time"

	dataMySQL "dbmsfinal/dataMySQL"
)

type OrderDAO struct{}

func (r *OrderDAO) GetOrderInfo(order_id int64) (*dataMySQL.Order, float64, error) {
	orderInfo := &dataMySQL.Order{}

	start := time.Now()

	row := DB.QueryRow("SELECT order_id, customer_id FROM Orders WHERE order_id = ?;", order_id)

	elapsed := time.Since(start).Seconds()

	err := row.Scan(&orderInfo.OrderID, &orderInfo.CustomerID)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get order info")
	}

	return orderInfo, elapsed, nil
}

func (r *OrderDAO) GetAllOrdersInfo() ([]*dataMySQL.Order, float64, error) {
	orders := make([]*dataMySQL.Order, 0)

	start := time.Now()

	rows, err := DB.Query("SELECT order_id, customer_id FROM Orders")

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orders info")
	}

	cnt := 0
	for rows.Next() {
		orderInfo := &dataMySQL.Order{}
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
