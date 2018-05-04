package queryMySQL

import (
	dataMySQL "dbmsfinal/dataMySQL"
	"errors"
	"fmt"
	"time"
)

type OrderdetailsDAO struct{}

func (r *OrderdetailsDAO) GetOrderdetailsOfOrderID(order_id int64) ([]*dataMySQL.Orderdetail, float64, error) {
	orderdetails := make([]*dataMySQL.Orderdetail, 0)

	start := time.Now()

	rows, err := DB.Query("SELECT o.order_id, o.customer_id, od.car_id, od.quantity_order FROM Orders o "+
		"JOIN Orderdetails od ON o.order_id = od.order_id "+
		"WHERE o.order_id = ?;", order_id)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetail of order_id")
	}

	cnt := 0
	for rows.Next() {
		orderdetail := &dataMySQL.Orderdetail{}
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

func (r *OrderdetailsDAO) GetOrderCardetailsOfOrderID(order_id int64) ([]*dataMySQL.OrderCarDetail, float64, error) {
	ordercardetails := make([]*dataMySQL.OrderCarDetail, 0)

	start := time.Now()

	rows, err := DB.Query("SELECT o.order_id, o.customer_id, od.car_id, c.car_model, od.quantity_order FROM Orders o "+
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
		ordercardetail := &dataMySQL.OrderCarDetail{}
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

func (r *OrderdetailsDAO) GetOrderCardetailsOfCustomerID(customer_id int64) ([]*dataMySQL.OrderCarDetail, float64, error) {
	ordercardetails := make([]*dataMySQL.OrderCarDetail, 0)

	start := time.Now()

	rows, err := DB.Query("SELECT o.order_id, o.customer_id, od.car_id, c.car_model, od.quantity_order FROM Orders o "+
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
		ordercardetail := &dataMySQL.OrderCarDetail{}
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

func (r *OrderdetailsDAO) InsertNewOrderdetail(insertOrderdetailsData *dataMySQL.OrderdetailData) (*dataMySQL.OrderdetailData, int64, error) {
	orderID := insertOrderdetailsData.OrderID
	carID := insertOrderdetailsData.CarID
	quantityOrder := insertOrderdetailsData.QuantityOrder

	start := time.Now()

	queryRequests, err1 := DB.Prepare("insert into Orderdetails(order_id, car_id, quantity_order) values (?, ?, ?);")
	///add vao bang orderdetais
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, int64(-1), errors.New("Fail to add new orderdetail")
	}

	_, err1 = queryRequests.Exec(orderID, carID, quantityOrder)

	elapsed := time.Since(start).Nanoseconds()

	if err1 != nil {
		fmt.Println(err1.Error())
		return nil, elapsed, errors.New("Fail to add new orderdetail")
	}

	return insertOrderdetailsData, elapsed, nil
}

func (r *OrderdetailsDAO) DeleteOrderdetail(order_id int64, car_id int64) (int64, error) { //TODO: DeleteOrderdetail
	start := time.Now()

	query, err := DB.Prepare("delete from Orderdetails where order_id= ? and car_id= ?;")
	if err != nil {
		fmt.Println(err.Error())
		return int64(-1), errors.New("Fail to delete Orderdetail")
	}
	_, err = query.Exec(order_id, car_id)
	if err != nil {
		fmt.Println(err.Error())
		return int64(-1), errors.New("Fail to delete Orderdetail")
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed, nil
}
