package queryMongo

import "gopkg.in/mgo.v2/bson"

type Car struct {
	ID           bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	CarID        int64         `bson:"car_id" json:"car_id"`
	CarModel     string        `bson:"car_model" json:"car_model"`
	CarMake      string        `bson:"car_make" json:"car_make"`
	CarModelYear int           `bson:"car_model_year" json:"car_model_year"`
}

type CarsDAO struct{}
