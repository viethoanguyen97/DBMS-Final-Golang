package queryMongo

import "gopkg.in/mgo.v2/bson"

type ODetail struct {
	CarID         int64 `bson:"car_id" json:"car_id"`
	QuantityOrder int64 `bson:"quantity_order" json:"quantity_order"`
}

type Orderdetail struct {
	ID      bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID int64         `bson:"order_id" json:"order_id"`
	Details []*ODetail    `bson:"details" json:"details"`
}

type OrderDetail struct {
	ID            bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID       int64         `bson:"order_id" json:"order_id"`
	CustomerID    int64         `bson:"customer_id" json:"customer_id"`
	CarID         int64         `bson:"car_id" json:"car_id"`
	CarModel      int64         `bson:"car_model" json:"car_model"`
	QuantityOrder int64         `bson:"quantity_order" json:"quantity_order"`
}

type OrderCardetail struct {
	ID      bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID int64         `bson:"order_id" json:"order_id"`
	Details []*ODetail    `bson:"details" json:"details"`
}

type OrderCarDetail struct {
	ID            bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID       int64         `bson:"order_id" json:"order_id"`
	CustomerID    int64         `bson:"customer_id" json:"customer_id"`
	CarID         int64         `bson:"car_id" json:"car_id"`
	CarModel      int64         `bson:"car_model" json:"car_model"`
	QuantityOrder int64         `bson:"quantity_order" json:"quantity_order"`
}

type OrderDetailsDAO struct{}
