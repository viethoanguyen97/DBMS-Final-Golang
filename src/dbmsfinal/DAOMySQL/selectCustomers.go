package queryMySQL

import (
	"errors"
	"fmt"
	"time"

	dataMySQL "dbmsfinal/dataMySQL"
)

type CustomerDAO struct{}

func (r *CustomerDAO) GetCustomerInfo(customer_id int64) (*dataMySQL.Customer, float64, error) {
	start := time.Now()

	customerInfo := &dataMySQL.Customer{}
	row := DB.QueryRow("SELECT customer_id, customer_name, customer_email, customer_address FROM Customers WHERE customer_id = ?;", customer_id)
	elapsed := time.Since(start).Seconds()

	err := row.Scan(&customerInfo.CustomerID, &customerInfo.CustomerName, &customerInfo.CustomerEmail, &customerInfo.CustomerAddress)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	return customerInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCustomersInfo() ([]*dataMySQL.Customer, float64, error) {
	customers := make([]*dataMySQL.Customer, 0)

	start := time.Now()
	rows, err := DB.Query("SELECT customer_id, customer_name, customer_email, customer_address FROM Customers")
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	cnt := 0
	for rows.Next() {
		customerInfo := &dataMySQL.Customer{}
		cnt++
		err := rows.Scan(&customerInfo.CustomerID, &customerInfo.CustomerName, &customerInfo.CustomerEmail, &customerInfo.CustomerAddress)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get all customers info")
		}

		customers = append(customers, customerInfo)
	}

	defer rows.Close()

	return customers, elapsed, nil
}
