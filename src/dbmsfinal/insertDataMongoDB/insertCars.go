package queryMongo

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	daoMongoDB "dbmsfinal/DAOMongoDB"
	dataMongoDB "dbmsfinal/dataMongoDB"

	"github.com/gocarina/gocsv"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getAllCars() []*dataMongoDB.Car {
	carsFile, err := os.Open("cars.json")
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()

	byteValue, _ := ioutil.ReadAll(carsFile)

	cars := []*dataMongoDB.Car{}

	err = json.Unmarshal(byteValue, &cars)
	if err != nil {
		panic(err)
	}

	return cars
}

func getAllCarsCSV() []*dataMongoDB.Car {
	carsFile, err := os.OpenFile("cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()

	cars := []*dataMongoDB.Car{}

	if err := gocsv.UnmarshalFile(carsFile, &cars); err != nil { // Load clients from file
		panic(err)
	}

	return cars
}

func insertCarsRowByRow() int64 {
	//cars := getAllCars()
	cars := getAllCarsCSV()
	daoMongoDB.Session.SetMode(mgo.Monotonic, true)

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMS-Final").C("Cars")
	//c.RemoveAll(nil)

	for _, car := range cars {
		//fmt.Println(*car, bson.NewObjectId())
		car.ID = bson.NewObjectId()
		err := c.Insert(car)
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start).Nanoseconds()

	return elapsed

}

func insertCarsBulk() int64 {
	cars := getAllCarsCSV()

	start := time.Now()
	c := daoMongoDB.Session.DB("DBMS-Final").C("Cars")
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

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}
