package queryMongo

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	ID              bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	CustomerID      int64         `bson:"customer_id" json:"customer_id"`
	CustomerName    string        `bson:"customer_name" json:"customer_name"`
	CustomerEmail   string        `bson:"customer_email" json:"customer_email"`
	CustomerAddress string        `bson:"customer_address" json:"customer_address"`
}
type CustomersDAO struct{}
