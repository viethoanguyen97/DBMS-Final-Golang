package queryMySQL

import (
	"os"

	gocsv "github.com/gocarina/gocsv"
)

type Orderdetail struct {
	OrderID       int64 `csv:"order_id"`
	CarID         int64 `csv:"car_id"`
	QuantityOrder int64 `csv:"quantity_order"`
}

func getAllOrderdetails() []*Orderdetail {
	orderdetailsFile, err := os.OpenFile("orderdetails.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	orderdetails := []*Orderdetail{}

	if err := gocsv.UnmarshalFile(orderdetailsFile, &orderdetails); err != nil { // Load clients from file
		panic(err)
	}

	return orderdetails
}

func insertOrderdetailsRowByRow() {
	orderdetails := getAllOrderdetails()
	sqlStr := "INSERT INTO Orderdetails(order_id, car_id, quantity_order) VALUES (?, ?, ?)"
	stmt, _ := db.Prepare(sqlStr)
	for _, orderdetail := range orderdetails {
		_, err := stmt.Exec(orderdetail.OrderID, orderdetail.CarID, orderdetail.QuantityOrder)
		if err != nil {
			panic(err)
		}
	}
}

func insertOrderdetailsBulk() {
	orderdetails := getAllOrderdetails()

	sqlStr := "INSERT INTO Orderdetails(order_id, car_id, quantity_order) VALUES "
	vals := []interface{}{}
	cnt := 0
	for _, orderdetail := range orderdetails {
		if cnt < 1000 {
			sqlStr += "(?, ?, ?),"
			cnt++
			vals = append(vals, orderdetail.OrderID, orderdetail.CarID, orderdetail.QuantityOrder)
		} else {
			sqlStr = sqlStr[0 : len(sqlStr)-1]
			stmt, _ := db.Prepare(sqlStr)
			_, err := stmt.Exec(vals...)
			if err != nil {
				panic(err)
			}
			sqlStr = "INSERT INTO Orderdetails(order_id, car_id, quantity_order) VALUES "
			cnt = 0
			vals = []interface{}{}
		}
	}

	if cnt < 1000 && cnt > 0 {
		sqlStr = sqlStr[0 : len(sqlStr)-1]
		stmt, _ := db.Prepare(sqlStr)
		_, err := stmt.Exec(vals...)
		if err != nil {
			panic(err)
		}
	}
}

