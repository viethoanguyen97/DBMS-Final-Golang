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
	query := Session.DB("DBMSFinal").C("Orders").Find(bson.M{"order_id": order_id})

	//Measure time execution
	err := query.One(orderInfo)
	elapsed := time.Since(start).Nanoseconds()
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
	query := Session.DB("DBMSFinal").C("Orders").Find(bson.M{})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.All(&orders)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orders info")
	}

	return orders, elapsed, nil
}

func (r *OrdersDAO) EditOrder(order_id int64, editOrderData *dataMongoDB.Order) (*dataMongoDB.Order, int64, error) {
	customer_id := editOrderData.CustomerID

	start := time.Now()

	err := Session.DB("DBMSFinal").C("Orders").Update(
		bson.M{
			"order_id": order_id,
		},
		bson.M{
			"customer_id": customer_id,
		},
	)

	elapsed := time.Since(start).Nanoseconds()

	editOrderData.OrderID = order_id

	if err != nil {
		return nil, elapsed, err
	}

	return editOrderData, elapsed, err
}

/*db.Cars.insert(
	{
		"car_id": 1001,
		"car_model": "Test insert 1",
		"car_make": "DCS",
		"car_model_year": 2018
	}
)*/
