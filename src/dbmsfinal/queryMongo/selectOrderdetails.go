package queryMongo

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (r *OrderDetailsDAO) GetOrderdetailsOfOrderID(order_id int64) ([]*OrderDetail, int64, error) {
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
	orderdetails := make([]*OrderDetail, 0)

	start := time.Now()

	collection := session.DB("DBMS-Final").C("Orders")
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
		bson.M{"$unwind": "$orderdetails.details"},
		bson.M{"$project": bson.M{
			"order_id":       "$order_id",
			"customer_id":    "$customer_id",
			"car_id":         "$orderdetails.details.car_id",
			"quantity_order": "$orderdetails.details.quantity_order",
		},
		},
	}
	pipe := collection.Pipe(pipeline)

	elapsed := time.Since(start).Nanoseconds()
	err := pipe.All(&orderdetails)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetails of order")
	}

	return orderdetails, elapsed, nil
}

func (r *OrderDetailsDAO) GetOrderCardetailsOfOrderID(order_id int64) ([]*OrderCarDetail, int64, error) {
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
	ordercardetails := make([]*OrderCarDetail, 0)

	start := time.Now()

	collection := session.DB("DBMS-Final").C("Orders")
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
	}
	pipe := collection.Pipe(pipeline)

	elapsed := time.Since(start).Nanoseconds()
	err := pipe.All(&ordercardetails)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetails of order")
	}

	return ordercardetails, elapsed, nil
}

func (r *OrderDetailsDAO) GetOrderCardetailsOfCustomerID(customer_id int64) ([]*OrderCarDetail, int64, error) {
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
	ordercardetails := make([]*OrderCarDetail, 0)

	start := time.Now()

	collection := session.DB("DBMS-Final").C("Orders")
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
	}
	pipe := collection.Pipe(pipeline)

	elapsed := time.Since(start).Nanoseconds()
	err := pipe.All(&ordercardetails)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get orderdetails of order")
	}

	return ordercardetails, elapsed, nil
}
