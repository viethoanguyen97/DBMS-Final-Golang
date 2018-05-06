package queryMongo

import (
	"errors"
	"fmt"
	"time"

	dataMongoDB "dbmsfinal/dataMongoDB"

	"gopkg.in/mgo.v2/bson"
)

type OrderDetailsDAO struct{}

func (r *OrderDetailsDAO) GetOrderdetailsOfOrderID(order_id int64) ([]*dataMongoDB.OrderDetail, float64, error) {
	/*db.Orders.aggregate([
	    {$match: {order_id: 150}},
	  	{$lookup: {
	  			from: "Orderdetails",
	  			localField: "order_id",
	  			foreignField: "order_id",
	  			as: "orderdetails"
	  		}
	  	},
	  	{$unwind: "$orderdetails"},
	  	{$unwind: "$orderdetails.details"},
	  	{$project: {
	  		"order_id": "$order_id",
	  		"customer_id": "$customer_id",
	  		"car_id": "$orderdetails.details.car_id",
	  		"quantity_order": "$orderdetails.details.quantity_order"
	  		}
	  	}
	  ]).pretty()*/
	orderdetails := make([]*dataMongoDB.OrderDetail, 0)

	start := time.Now()

	collection := Session.DB("DBMSFinal").C("Orders")
	/*pipeline := []bson.M{
		bson.M{"$match": bson.M{"order_id": order_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$unwind": "$orderdetails.details"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$orderdetails.details.car_id",
			"quantity_order": "$orderdetails.details.quantity_order",
		},
		},
	}*/

	/*db.Orders.aggregate([
		{$match: {order_id: 150}},
		{$lookup: {
				from: "Orderdetails",
				localField: "order_id",
				foreignField: "order_id",
				as: "orderdetails"
			}
		},
		{$unwind: "$orderdetails"},
		{$project: {
			"order_id": "$order_id",
			"customer_id": "$customer_id",
			"car_id": "$orderdetails.car_id",
			"quantity_order": "$orderdetails.quantity_order"
			}
		}
	]).pretty()*/
	pipeline := []bson.M{
		bson.M{"$match": bson.M{"order_id": order_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$orderdetails.car_id",
			"quantity_order": "$orderdetails.quantity_order",
		},
		},
	}
	pipe := collection.Pipe(pipeline)

	err := pipe.All(&orderdetails)

	elapsed := time.Since(start).Seconds()
	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetails of order")
	}

	return orderdetails, elapsed, nil
}

func (r *OrderDetailsDAO) GetOrderCardetailsOfOrderID(order_id int64) ([]*dataMongoDB.OrderCarDetail, float64, error) {
	/*db.Orders.aggregate([
	    {$match: {order_id: 150}},

	  	{$lookup: {
	        from: "Orderdetails",
	        localField: "order_id",
	        foreignField: "order_id",
	        as: "orderdetails"
	      }
	  	},
	  	{$unwind: "$orderdetails"},
	  	{$unwind: "$orderdetails.details"},
	    {$lookup: {
	        from: "Cars",
	        localField: "orderdetails.details.car_id",
	        foreignField: "car_id",
	        as: "car_info"
	      }
	  	},
	  	{$unwind: "$car_info"},
	  	{$project: {
	  		"order_id": "$order_id",
	  		"customer_id": "$customer_id",
	  		"car_id": "$car_info.car_id",
	  		"car_model": "$car_info.car_model",
	  		"quantity_order": "$details.quantity_order"
	  		}
	  	}
	  ]).pretty()*/
	ordercardetails := make([]*dataMongoDB.OrderCarDetail, 0)

	start := time.Now()

	collection := Session.DB("DBMSFinal").C("Orders")
	/*pipeline := []bson.M{
		bson.M{"$match": bson.M{"order_id": order_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$unwind": "$orderdetails.details"},
		bson.M{"$lookup": bson.M{
			"from":         "Cars",
			"localField":   "orderdetails.details.car_id",
			"foreignField": "car_id",
			"as":           "car_info",
		},
		},
		bson.M{"$unwind": "$car_info"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$car_info.car_id",
			"car_model":      "$car_info.car_model",
			"quantity_order": "$details.quantity_order",
		},
		},
	}*/

	/*db.Orders.aggregate([
		{$match: {order_id: 150}},

		{$lookup: {
				from: "Orderdetails",
				localField: "order_id",
				foreignField: "order_id",
				as: "orderdetails"
			}
		},
		{$unwind: "$orderdetails"},
		{$lookup: {
				from: "Cars",
				localField: "orderdetails.car_id",
				foreignField: "car_id",
				as: "car_info"
			}
		},
		{$unwind: "$car_info"},
		{$project: {
			"order_id": "$order_id",
			"customer_id": "$customer_id",
			"car_id": "$car_info.car_id",
			"car_model": "$car_info.car_model",
			"quantity_order": "$orderdetails.quantity_order"
			}
		}
	]).pretty()*/

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"order_id": order_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$lookup": bson.M{
			"from":         "Cars",
			"localField":   "orderdetails.car_id",
			"foreignField": "car_id",
			"as":           "car_info",
		},
		},
		bson.M{"$unwind": "$car_info"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$car_info.car_id",
			"car_model":      "$car_info.car_model",
			"quantity_order": "$orderdetails.quantity_order",
		},
		},
	}
	pipe := collection.Pipe(pipeline)

	fmt.Println(time.Since(start).Seconds())
	err := pipe.All(&ordercardetails)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetails of order")
	}

	return ordercardetails, elapsed, nil
}

func (r *OrderDetailsDAO) GetOrderCardetailsOfCustomerID(customer_id int64) ([]*dataMongoDB.OrderCarDetail, float64, error) {
	/*db.Orders.aggregate([
	  {$match: {customer_id: 150}},

	  {$lookup: {
	      from: "Orderdetails",
	      localField: "order_id",
	      foreignField: "order_id",
	      as: "orderdetails"
	    }
	  },
	  {$unwind: "$orderdetails"},
	  {$unwind: "$orderdetails.details"},
	  {$lookup: {
	      from: "Cars",
	      localField: "orderdetails.details.car_id",
	      foreignField: "car_id",
	      as: "car_info"
	    }
	  },
	  {$unwind: "$car_info"},
	  {$project: {
	    "order_id": "$order_id",
	    "customer_id": "$customer_id",
	    "car_id": "$car_info.car_id",
	    "car_model": "$car_info.car_model",
	    "quantity_order": "$details.quantity_order"
	    }
	  }
	]).pretty()*/
	ordercardetails := make([]*dataMongoDB.OrderCarDetail, 0)

	start := time.Now()

	collection := Session.DB("DBMSFinal").C("Orders")
	/*pipeline := []bson.M{
		bson.M{"$match": bson.M{"customer_id": customer_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$unwind": "$orderdetails.details"},
		bson.M{"$lookup": bson.M{
			"from":         "Cars",
			"localField":   "orderdetails.details.car_id",
			"foreignField": "car_id",
			"as":           "car_info",
		},
		},
		bson.M{"$unwind": "$car_info"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$car_info.car_id",
			"car_model":      "$car_info.car_model",
			"quantity_order": "$details.quantity_order",
		},
		},
	}*/

	/*db.Orders.aggregate([
		{$match: {customer_id: 150}},

		{$lookup: {
				from: "Orderdetails",
				localField: "order_id",
				foreignField: "order_id",
				as: "orderdetails"
			}
		},
		{$unwind: "$orderdetails"},
		{$lookup: {
				from: "Cars",
				localField: "orderdetails.car_id",
				foreignField: "car_id",
				as: "car_info"
			}
		},
		{$unwind: "$car_info"},
		{$project: {
			"order_id": "$order_id",
			"customer_id": "$customer_id",
			"car_id": "$car_info.car_id",
			"car_model": "$car_info.car_model",
			"quantity_order": "$orderdetails.quantity_order"
			}
		}
	]).pretty()*/

	pipeline := []bson.M{
		bson.M{"$match": bson.M{"customer_id": customer_id}},
		bson.M{"$lookup": bson.M{
			"from":         "Orderdetails",
			"localField":   "order_id",
			"foreignField": "order_id",
			"as":           "orderdetails",
		},
		},
		bson.M{"$unwind": "$orderdetails"},
		bson.M{"$lookup": bson.M{
			"from":         "Cars",
			"localField":   "orderdetails.car_id",
			"foreignField": "car_id",
			"as":           "car_info",
		},
		},
		bson.M{"$unwind": "$car_info"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$car_info.car_id",
			"car_model":      "$car_info.car_model",
			"quantity_order": "$orderdetails.quantity_order",
		},
		},
	}
	pipe := collection.Pipe(pipeline)
	fmt.Println(time.Since(start).Seconds())
	err := pipe.All(&ordercardetails)
	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all orderdetails of customer")
	}

	return ordercardetails, elapsed, nil
}

func (*OrderDetailsDAO) InsertNewOrderdetail(insertNewOrderdetailData *dataMongoDB.OrderdetailCSV) (*dataMongoDB.OrderdetailCSV, float64, error) {
	insertNewOrderdetailData.ID = bson.NewObjectId()
	start := time.Now()

	err := Session.DB("DBMSFinal").C("Orderdetails").Insert(insertNewOrderdetailData)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		panic(err)
		return nil, elapsed, errors.New("Fail to insert new orderdetail")
	}

	return insertNewOrderdetailData, elapsed, nil
}

func (*OrderDetailsDAO) DeleteOrderdetail(order_id int64, car_id int64) (float64, error) {
	start := time.Now()

	err := Session.DB("DBMSFinal").C("Orderdetails").Remove(
		bson.M{
			"order_id": order_id,
			"car_id":   car_id,
		},
	)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		panic(err)
		return elapsed, errors.New("Fail to delete orderdetail")
	}

	return elapsed, nil
}
