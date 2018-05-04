package queryMongo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	daoMongoDB "dbmsfinal/DAOMongoDB"
	dataMongoDB "dbmsfinal/dataMongoDB"

	"github.com/gocarina/gocsv"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAllOrderdetails() []*dataMongoDB.Orderdetail {
	orderdetailsFile, err := os.Open("orderwithdetails.json")
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	byteValue, _ := ioutil.ReadAll(orderdetailsFile)

	orderdetails := []*dataMongoDB.Orderdetail{}

	err = json.Unmarshal(byteValue, &orderdetails)
	if err != nil {
		panic(err)
	}

	return orderdetails
}

func getAllOrderdetailsCSV() []*dataMongoDB.Orderdetail {
	orderdetailsFile, err := os.OpenFile("orderdetails_final.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	orderdetailscsv := []*dataMongoDB.OrderdetailCSV{}

	if err := gocsv.UnmarshalFile(orderdetailsFile, &orderdetailscsv); err != nil { // Load clients from file
		panic(err)
	}

	orderdetails := make([]*dataMongoDB.Orderdetail, 0)
	details := make([]*dataMongoDB.ODetail, 0)

	currentOrderID := int64(1)
	for _, orderdetailcsv := range orderdetailscsv {
		orderID := orderdetailcsv.OrderID
		carID := orderdetailcsv.CarID
		quantityOrder := orderdetailcsv.QuantityOrder
		detail := &dataMongoDB.ODetail{
			CarID:         carID,
			QuantityOrder: quantityOrder,
		}

		if orderID != currentOrderID {
			orderdetail := &dataMongoDB.Orderdetail{
				OrderID: currentOrderID,
				Details: details,
			}

			orderdetails = append(orderdetails, orderdetail)
			details = make([]*dataMongoDB.ODetail, 0)
			details = append(details, detail)
		} else {
			details = append(details, detail)
		}

		currentOrderID = orderID
	}

	last_order_detail := &dataMongoDB.Orderdetail{
		OrderID: currentOrderID,
		Details: details,
	}

	orderdetails = append(orderdetails, last_order_detail)

	return orderdetails
}

func getAllOrderdetailsCSVOriginal() []*dataMongoDB.OrderdetailCSV {
	orderdetailsFile, err := os.OpenFile("orderdetails_final.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer orderdetailsFile.Close()

	orderdetailscsv := []*dataMongoDB.OrderdetailCSV{}

	if err := gocsv.UnmarshalFile(orderdetailsFile, &orderdetailscsv); err != nil { // Load clients from file
		panic(err)
	}

	return orderdetailscsv
}

func InsertOrderdetailsRowByRow() int64 {
	//orderdetails := getAllOrderdetails()
	orderdetails := getAllOrderdetailsCSVOriginal()

	fmt.Println("ok")
	daoMongoDB.Session.SetMode(mgo.Monotonic, true)

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMSFinal").C("Orderdetails")
	//c.RemoveAll(nil)

	for _, orderdetail := range orderdetails {
		orderdetail.ID = bson.NewObjectId()
		err := c.Insert(orderdetail)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}

func InsertOrderdetailsBulk() int64 {
	//orderdetails := getAllOrderdetails()
	orderdetails := getAllOrderdetailsCSVOriginal()

	fmt.Println("OK")
	start := time.Now()

	c := daoMongoDB.Session.DB("DBMSFinal").C("Orderdetails")
	//c.RemoveAll(nil)
	bulk := c.Bulk()
	cnt := 0

	for _, orderdetail := range orderdetails {
		//fmt.Println(*car, bson.NewObjectId())
		orderdetail.ID = bson.NewObjectId()
		bulk.Insert(orderdetail)
		if cnt == 100000 {
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
