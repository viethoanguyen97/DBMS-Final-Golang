package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	daoMongoDB "dbmsfinal/DAOMongoDB"
	dataMongoDB "dbmsfinal/dataMongoDB"

	"github.com/gocarina/gocsv"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAllCustomers() []*dataMongoDB.Customer {
	customersFile, err := os.Open("customers.json")
	if err != nil {
		panic(err)
	}
	defer customersFile.Close()

	byteValue, _ := ioutil.ReadAll(customersFile)

	customers := []*dataMongoDB.Customer{}

	err = json.Unmarshal(byteValue, &customers)
	if err != nil {
		panic(err)
	}

	return customers
}

func getAllCustomersCSV() []*dataMongoDB.Customer {
	customersFile, err := os.OpenFile("customers_final.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer customersFile.Close()

	customers := []*dataMongoDB.Customer{}

	if err := gocsv.UnmarshalFile(customersFile, &customers); err != nil { // Load clients from file
		panic(err)
	}

	return customers
}

func insertCustomersRowByRow() int64 {
	//customers := getAllCustomers()
	customers := getAllCustomersCSV()
	daoMongoDB.Session.SetMode(mgo.Monotonic, true)

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMSFinal").C("Customers")
	//c.RemoveAll(nil)

	for _, customer := range customers {
		//fmt.Println(*car, bson.NewObjectId())
		customer.ID = bson.NewObjectId()
		err := c.Insert(customer)
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}

func insertCustomersBulk() int64 {
	//customers := getAllCustomers()
	customers := getAllCustomersCSV()
	start := time.Now()

	c := daoMongoDB.Session.DB("DBMSFinal").C("Customers")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	cnt := 0
	for _, customer := range customers {
		cnt++
		customer.ID = bson.NewObjectId()
		bulk.Insert(customer)

		if cnt == 5000 {
			cnt = 0
			_, err := bulk.Run()

			if err != nil {
				panic(err)
			}

			bulk = c.Bulk()
		}
	}

	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}
