package queryMySQL

import (
	"os"
	"time"

	daoMySQL "dbmsfinal/DAOMySQL"
	dataMySQL "dbmsfinal/dataMySQL"

	gocsv "github.com/gocarina/gocsv"
)

func getAllCars() []*dataMySQL.Car {
	carsFile, err := os.OpenFile("cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()

	cars := []*dataMySQL.Car{}

	if err := gocsv.UnmarshalFile(carsFile, &cars); err != nil { // Load clients from file
		panic(err)
	}

	return cars
}

func insertCarsRowByRow() int64 {
	cars := getAllCars()

	start := time.Now()

	sqlStr := "INSERT INTO Cars(car_id, car_model, car_make, car_model_year) VALUES (?, ?, ?, ?)"
	stmt, _ := daoMySQL.DB.Prepare(sqlStr)
	for _, car := range cars {
		_, err := stmt.Exec(car.CarID, car.CarModel, car.CarMake, car.CarModelYear)
		if err != nil {
			panic(err)
		}
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}

func insertCarsBulk() int64 {
	cars := getAllCars()

	start := time.Now()
	sqlStr := "INSERT INTO Cars(car_id, car_model, car_make, car_model_year) VALUES "
	vals := []interface{}{}
	for _, car := range cars {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, car.CarID, car.CarModel, car.CarMake, car.CarModelYear)

	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, _ := daoMySQL.DB.Prepare(sqlStr)
	_, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start).Nanoseconds()

	return elapsed
}
