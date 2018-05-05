package queryMongo

import (
	dataMongoDB "dbmsfinal/dataMongoDB"
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CustomersDAO struct{}

func (r *CustomersDAO) GetCustomerInfo(customer_id int64) (*dataMongoDB.Customer, float64, error) {
	customerInfo := &dataMongoDB.Customer{}

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Customers").Find(bson.M{"customer_id": customer_id})
	elapsed := time.Since(start).Seconds()
	//Measure time execution

	err := query.One(&customerInfo)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get customer info")
	}

	return customerInfo, elapsed, nil
}

func (r *CustomersDAO) GetAllCustomersInfo() (int, float64, error) { //([]*dataMongoDB.Customer, float64, error) {
	//customers := make([]*dataMongoDB.Customer, 0)

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Customers").Find(bson.M{})
	//Measure time execution

	//	err := query.All(&customers)
	cnt, err := query.Count()
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return 0, elapsed, err
		//return nil, elapsed, errors.New("Fail to get all customers info")
	}

	return cnt, elapsed, nil
	//	return customers, elapsed, nil
}
