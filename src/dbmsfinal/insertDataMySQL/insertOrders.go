package queryMySQL

import (
	"os"
	"time"

	daoMySQL "dbmsfinal/DAOMySQL"
	dataMySQL "dbmsfinal/dataMySQL"

	gocsv "github.com/gocarina/gocsv"
)

func getAllOrders() []*dataMySQL.Order {
	ordersFile, err := os.OpenFile("orders.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ordersFile.Close()

	orders := []*dataMySQL.Order{}

	if err := gocsv.UnmarshalFile(ordersFile, &orders); err != nil { // Load clients from file
		panic(err)
	}

	return orders
}

func insertOrdersRowByRow() int64 {
	orders := getAllOrders()

	start := time.Now()

	sqlStr := "INSERT INTO Orders(order_id, customer_id) VALUES (?, ?)"
	stmt, _ := daoMySQL.DB.Prepare(sqlStr)
	for _, order := range orders {
		_, err := stmt.Exec(order.OrderID, order.CustomerID)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()
	return elapsed
}

func insertOrdersBulk() int64 {
	orders := getAllOrders()

	start := time.Now()
	sqlStr := "INSERT INTO Orders(order_id, customer_id) VALUES "
	vals := []interface{}{}
	for _, order := range orders {
		sqlStr += "(?, ?),"
		vals = append(vals, order.OrderID, order.CustomerID)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, _ := daoMySQL.DB.Prepare(sqlStr)
	_, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start).Nanoseconds()
	return elapsed
}
