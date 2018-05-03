package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAllCustomers() []*Customer {
	customersFile, err := os.Open("customers.json")
	if err != nil {
		panic(err)
	}
	defer customersFile.Close()

	byteValue, _ := ioutil.ReadAll(customersFile)

	customers := []*Customer{}

	err = json.Unmarshal(byteValue, &customers)
	if err != nil {
		panic(err)
	}

	return customers
}

func insertCustomersRowByRow() {
	customers := getAllCustomers()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("DBMS-Final").C("Customers")
	//c.RemoveAll(nil)

	for _, customer := range customers {
		//fmt.Println(*car, bson.NewObjectId())
		customer.ID = bson.NewObjectId()
		err := c.Insert(customer)
		if err != nil {
			panic(err)
		}
	}
}

func insertCustomersBulk() {
	customers := getAllCustomers()

	c := session.DB("DBMS-Final").C("Customers")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	for _, customer := range customers {
		//fmt.Println(*car, bson.NewObjectId())
		customer.ID = bson.NewObjectId()
		bulk.Insert(customer)
	}
	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}
}
