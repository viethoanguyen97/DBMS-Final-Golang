package queryMySQL

import (
	"os"
	"time"

	daoMySQL "dbmsfinal/DAOMySQL"
	dataMySQL "dbmsfinal/dataMySQL"

	gocsv "github.com/gocarina/gocsv"
)

func getAllOrderdetails() []*dataMySQL.Orderdetail {
	orderdetailsFile, err := os.OpenFile("orderdetails.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	orderdetails := []*dataMySQL.Orderdetail{}

	if err := gocsv.UnmarshalFile(orderdetailsFile, &orderdetails); err != nil { // Load clients from file
		panic(err)
	}

	return orderdetails
}

func insertOrderdetailsRowByRow() int64 {
	orderdetails := getAllOrderdetails()

	start := time.Now()

	sqlStr := "INSERT INTO Orderdetails(order_id, car_id, quantity_order) VALUES (?, ?, ?)"
	stmt, _ := daoMySQL.DB.Prepare(sqlStr)
	for _, orderdetail := range orderdetails {
		_, err := stmt.Exec(orderdetail.OrderID, orderdetail.CarID, orderdetail.QuantityOrder)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()
	return elapsed
}

func insertOrderdetailsBulk() int64 {
	orderdetails := getAllOrderdetails()

	start := time.Now()

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
			stmt, _ := daoMySQL.DB.Prepare(sqlStr)
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
		stmt, _ := daoMySQL.DB.Prepare(sqlStr)
		_, err := stmt.Exec(vals...)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()
	return elapsed
}
