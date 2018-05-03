package queryMySQL

import (
	"errors"
	"fmt"
	"time"
)

func (r *CustomerDAO) GetCustomerInfo(customer_id int64) (*Customer, float64, error) {
	start := time.Now()

	customerInfo := &Customer{}

	row := db.QueryRow("SELECT customer_id, customer_name, customer_email, customer_address FROM Customers WHERE customer_id = ?;", customer_id)

	elapsed := time.Since(start).Seconds()

	err := row.Scan(&customerInfo.CustomerID, &customerInfo.CustomerName, &customerInfo.CustomerEmail, &customerInfo.CustomerAddress)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	return customerInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCustomersInfo() ([]*Customer, float64, error) {
	customers := make([]*Customer, 0)
	start := time.Now()

	rows, err := db.Query("SELECT customer_id, customer_name, customer_email, customer_address FROM Customers")

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	cnt := 0
	for rows.Next() {
		customerInfo := &Customer{}
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
