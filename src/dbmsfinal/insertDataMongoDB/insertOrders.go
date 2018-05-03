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

func getAllOrders() []*dataMongoDB.Order {
	ordersFile, err := os.Open("orders.json")
	if err != nil {
		panic(err)
	}
	defer ordersFile.Close()

	byteValue, _ := ioutil.ReadAll(ordersFile)

	orders := []*dataMongoDB.Order{}

	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		panic(err)
	}

	return orders
}

func getAllOrdersCSV() []*dataMongoDB.Order {
	ordersFile, err := os.OpenFile("orders.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer ordersFile.Close()

	orders := []*dataMongoDB.Order{}

	if err := gocsv.UnmarshalFile(ordersFile, &orders); err != nil { // Load clients from file
		panic(err)
	}

	return orders
}

func insertOrdersRowByRow() int64 {
	//orders := getAllOrders()
	orders := getAllOrdersCSV()

	daoMongoDB.Session.SetMode(mgo.Monotonic, true)

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMS-Final").C("Orders")
	//c.RemoveAll(nil)

	for _, order := range orders {
		//fmt.Println(*car, bson.NewObjectId())
		order.ID = bson.NewObjectId()
		err := c.Insert(order)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()
	return elapsed
}

func insertOrdersBulk() int64 {
	//orders := getAllOrders()
	orders := getAllOrdersCSV()

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMS-Final").C("Orders")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	for _, order := range orders {
		//fmt.Println(*car, bson.NewObjectId())
		order.ID = bson.NewObjectId()
		bulk.Insert(order)
	}
	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}
