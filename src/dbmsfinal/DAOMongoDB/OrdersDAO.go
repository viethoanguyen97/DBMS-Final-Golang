package queryMongo

import (
	"errors"
	"fmt"
	"time"

	dataMongoDB "dbmsfinal/dataMongoDB"

	"gopkg.in/mgo.v2/bson"
)

type OrdersDAO struct{}

func (r *OrdersDAO) GetOrderInfo(order_id int64) (*dataMongoDB.Order, int64, error) {
	orderInfo := &dataMongoDB.Order{}

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMS-Final").C("Orders").Find(bson.M{"order_id": order_id})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.One(&orderInfo)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get order info")
	}

	return orderInfo, elapsed, nil
}

func (r *OrdersDAO) GetAllOrdersInfo() ([]*dataMongoDB.Order, int64, error) {
	orders := make([]*dataMongoDB.Order, 0)

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMS-Final").C("Orders").Find(bson.M{})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.All(&orders)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orders info")
	}

	return orders, elapsed, nil
}
