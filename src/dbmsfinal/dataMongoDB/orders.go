package queryMongo

import "gopkg.in/mgo.v2/bson"

type Order struct {
	ID         bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	OrderID    int64         `bson:"order_id" csv:"order_id" json:"order_id" omit:empty`
	CustomerID int64         `bson:"customer_id" csv:"customer_id" json:"customer_id"`
}
