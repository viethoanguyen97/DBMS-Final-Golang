package queryMongo

import "gopkg.in/mgo.v2/bson"

type Order struct {
	ID         bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID    int64         `bson:"order_id" json:"order_id"`
	CustomerID int64         `bson:"customer_id" json:"customer_id"`
}

type OrdersDAO struct{}
