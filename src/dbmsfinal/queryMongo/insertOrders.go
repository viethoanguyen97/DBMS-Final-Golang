package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID         bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID    int64         `bson:"order_id" json:"order_id"`
	CustomerID int64         `bson:"customer_id" json:"customer_id"`
}

func getAllOrders() []*Order {
	ordersFile, err := os.Open("orders.json")
	if err != nil {
		panic(err)
	}
	defer ordersFile.Close()

	byteValue, _ := ioutil.ReadAll(ordersFile)

	orders := []*Order{}

	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		panic(err)
	}

	return orders
}

func insertOrdersRowByRow() {
	orders := getAllOrders()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("DBMS-Final").C("Orders")
	//c.RemoveAll(nil)

	for _, order := range orders {
		//fmt.Println(*car, bson.NewObjectId())
		order.ID = bson.NewObjectId()
		err := c.Insert(order)
		if err != nil {
			panic(err)
		}
	}
}

func insertOrdersBulk() {
	orders := getAllOrders()

	c := session.DB("DBMS-Final").C("Orders")
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
}

