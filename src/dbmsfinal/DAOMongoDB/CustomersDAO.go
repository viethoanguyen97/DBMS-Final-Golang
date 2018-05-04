package queryMongo

import (
	dataMongoDB "dbmsfinal/dataMongoDB"
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CustomersDAO struct{}

func (r *CustomersDAO) GetCustomerInfo(customer_id int64) (*dataMongoDB.Customer, int64, error) {
	customerInfo := &dataMongoDB.Customer{}

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Customers").Find(bson.M{"customer_id": customer_id})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.One(&customerInfo)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	return customerInfo, elapsed, nil
}

func (r *CustomersDAO) GetAllCustomersInfo() ([]*dataMongoDB.Customer, int64, error) {
	customers := make([]*dataMongoDB.Customer, 0)

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Customers").Find(bson.M{})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.All(&customers)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all customers info")
	}

	return customers, elapsed, nil
}
