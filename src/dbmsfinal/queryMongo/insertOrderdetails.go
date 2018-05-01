package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
)

type ODetail struct {
	CarID         int64 `bson:"car_id" json:"car_id"`
	QuantityOrder int64 `bson:"quantity_order" json:"quantity_order"`
}

type Orderdetail struct {
	//ID      bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID int64      `bson:"order_id" json:"order_id"`
	Details []*ODetail `bson:"details" json:"details"`
}

func getAllOrderdetails() []*Orderdetail {
	orderdetailsFile, err := os.Open("orderwithdetails.json")
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	byteValue, _ := ioutil.ReadAll(orderdetailsFile)

	orderdetails := []*Orderdetail{}

	err = json.Unmarshal(byteValue, &orderdetails)
	if err != nil {
		panic(err)
	}

	return orderdetails
}

func insertOrderdetailsRowByRow() {
	orderdetails := getAllOrderdetails()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("DBMS-Final").C("Orderdetails")
	//c.RemoveAll(nil)

	for _, orderdetail := range orderdetails {
		//fmt.Println(*car, bson.NewObjectId())
		//orderdetail.ID = bson.NewObjectId()
		err := c.Insert(orderdetail)
		if err != nil {
			panic(err)
		}
	}
}

func insertOrderdetailsBulk() {
	orderdetails := getAllOrderdetails()

	c := session.DB("DBMS-Final").C("Orderdetails")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	for _, orderdetail := range orderdetails {
		//fmt.Println(*car, bson.NewObjectId())
		//	orderdetail.ID = bson.NewObjectId()
		bulk.Insert(orderdetail)
	}
	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}
}

