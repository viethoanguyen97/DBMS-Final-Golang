package queryMongo

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (r *OrdersDAO) GetOrderInfo(order_id int64) (*Order, int64, error) {
	orderInfo := &Order{}

	//Measure time execution
	start := time.Now()
	query := session.DB("DBMS-Final").C("Orders").Find(bson.M{"order_id": order_id})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.One(&orderInfo)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get order info")
	}

	return orderInfo, elapsed, nil
}

func (r *OrdersDAO) GetAllOrdersInfo() ([]*Order, int64, error) {
	orders := make([]*Order, 0)

	//Measure time execution
	start := time.Now()
	query := session.DB("DBMS-Final").C("Orders").Find(bson.M{})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.All(&orders)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orders info")
	}

	return orders, elapsed, nil
}
