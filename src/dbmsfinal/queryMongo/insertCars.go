package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Car struct {
	ID           bson.ObjectId `bson:"_id" json:"_id, omitempty"`
	CarID        int64         `bson:"car_id" json:"car_id"`
	CarModel     string        `bson:"car_model" json:"car_model"`
	CarMake      string        `bson:"car_make" json:"car_make"`
	CarModelYear int           `bson:"car_model_year" json:"car_model_year"`
}

func getAllCars() []*Car {
	carsFile, err := os.Open("cars.json")
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()

	byteValue, _ := ioutil.ReadAll(carsFile)

	cars := []*Car{}

	err = json.Unmarshal(byteValue, &cars)
	if err != nil {
		panic(err)
	}

	return cars
}

func insertCarsRowByRow() {
	cars := getAllCars()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("DBMS-Final").C("Cars")
	c.RemoveAll(nil)

	for _, car := range cars {
		//fmt.Println(*car, bson.NewObjectId())
		car.ID = bson.NewObjectId()
		err := c.Insert(car)
		if err != nil {
			panic(err)
		}
	}
}

func insertCarsBulk() {
	cars := getAllCars()

	c := session.DB("DBMS-Final").C("Cars")
	//c.RemoveAll(nil)
	bulk := c.Bulk()

	for _, car := range cars {
		//fmt.Println(*car, bson.NewObjectId())
		car.ID = bson.NewObjectId()
		bulk.Insert(car)
	}
	_, err := bulk.Run()

	if err != nil {
		panic(err)
	}
}

