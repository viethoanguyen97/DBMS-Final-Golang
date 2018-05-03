package queryMongo

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func (r *CarsDAO) GetCarInfo(car_id int64) (*Car, int64, error) {
	carInfo := &Car{}

	//Measure time execution
	start := time.Now()
	query := session.DB("DBMS-Final").C("Cars").Find(bson.M{"car_id": car_id})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.One(&carInfo)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get car info")
	}

	return carInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCarsInfo() ([]*Car, int64, error) {
	cars := make([]*Car, 0)

	//Measure time execution
	start := time.Now()
	query := session.DB("DBMS-Final").C("Cars").Find(bson.M{})
	elapsed := time.Since(start).Nanoseconds()
	//Measure time execution

	err := query.All(&cars)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all cars info")
	}

	return cars, elapsed, nil
}
