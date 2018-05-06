package queryMongo

import (
	"errors"
	"fmt"
	"time"

	dataMongoDB "dbmsfinal/dataMongoDB"

	"gopkg.in/mgo.v2/bson"
)

type OrdersDAO struct{}

type Count struct {
	Count int64 `bson:"count" json:"count"`
}

func (r *OrdersDAO) GetOrderInfo(order_id int64) (*dataMongoDB.Order, float64, error) {
	orderInfo := &dataMongoDB.Order{}

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Orders").Find(bson.M{"order_id": order_id})

	//Measure time execution
	err := query.One(orderInfo)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get order info")
	}

	return orderInfo, elapsed, nil
}

func (r *OrdersDAO) GetAllOrdersInfo() (int64, float64, error) { // ([]*dataMongoDB.Order, int64, error) {
	//orders := make([]*dataMongoDB.Order, 0)

	//Measure time execution
	start := time.Now()
	//query := Session.DB("DBMSFinal").C("Orders").Find(bson.M{})

	//Measure time execution
	//err := query.All(&orders)
	//cnt, err := query.Count()
	//db.orderdetails.aggregate([{ $group: { _id: null, count: { $sum: 1 } } }])
	collection := Session.DB("DBMSFinal").C("Orders")
	pipeline := []bson.M{
		bson.M{"$group": bson.M{
			"_id": bson.M{},
			"count": bson.M{
				"$sum": 1,
			},
		},
		},
	}

	count := &Count{}
	pipe := collection.Pipe(pipeline)
	err := pipe.One(count)

	fmt.Println(count)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		//return nil, elapsed, errors.New("Fail to get all orders info")
		return 0, elapsed, err
	}

	return count.Count, elapsed, nil //int64(cnt), elapsed, nil
	//return orders, elapsed, nil
}

func (r *OrdersDAO) EditOrder(order_id int64, editOrderData *dataMongoDB.Order) (*dataMongoDB.Order, float64, error) {
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

	elapsed := time.Since(start).Seconds()

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
