package queryMongo

import (
	dataMongoDB "dbmsfinal/dataMongoDB"
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CarsDAO struct{}

func (r *CarsDAO) GetCarInfo(car_id int64) (*dataMongoDB.Car, float64, error) {
	carInfo := &dataMongoDB.Car{}

	//Measure time execution
	start := time.Now()
	query := Session.DB("DBMSFinal").C("Cars").Find(bson.M{"car_id": car_id})

	//Measure time execution

	err := query.One(&carInfo)
	elapsed := time.Since(start).Seconds()
	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get car info")
	}

	return carInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCarsInfo() (int64, float64, error) { //([]*dataMongoDB.Car, int64, error) {
	//cars := make([]*dataMongoDB.Car, 0)

	//Measure time execution
	start := time.Now()
	//	query := Session.DB("DBMSFinal").C("Cars").Find(bson.M{})

	//Measure time execution

	//	err := query.All(&cars)
	//cnt, err := query.Count()
	collection := Session.DB("DBMSFinal").C("Cars")
	pipeline := []bson.M{
		bson.M{"$group": bson.M{
			"_id": bson.M{},
			"count": bson.M{
				"$sum": 1,
			},
		},
		},
	}

	count := &Count{}
	pipe := collection.Pipe(pipeline)
	err := pipe.One(count)

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return 0, elapsed, err
		//return nil, elapsed, errors.New("Fail to get all cars info")
	}

	return count.Count, elapsed, nil
	//	return cars, elapsed, nil
}
