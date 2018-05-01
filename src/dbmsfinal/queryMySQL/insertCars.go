package queryMySQL

import (
	"os"

	gocsv "github.com/gocarina/gocsv"
)

type Car struct {
	CarID        int64  `csv:"car_id"`
	CarModel     string `csv:"car_model"`
	CarMake      string `csv:"car_make"`
	CarModelYear int    `csv:"car_model_year"`
}

func getAllCars() []*Car {
	carsFile, err := os.OpenFile("cars.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer carsFile.Close()

	cars := []*Car{}

	if err := gocsv.UnmarshalFile(carsFile, &cars); err != nil { // Load clients from file
		panic(err)
	}

	return cars
}

func insertCarsRowByRow() {
	cars := getAllCars()
	sqlStr := "INSERT INTO Cars(car_id, car_model, car_make, car_model_year) VALUES (?, ?, ?, ?)"
	stmt, _ := db.Prepare(sqlStr)
	for _, car := range cars {
		_, err := stmt.Exec(car.CarID, car.CarModel, car.CarMake, car.CarModelYear)
		if err != nil {
			panic(err)
		}
	}
}

func insertCarsBulk() {
	cars := getAllCars()

	sqlStr := "INSERT INTO Cars(car_id, car_model, car_make, car_model_year) VALUES "
	vals := []interface{}{}
	for _, car := range cars {
		sqlStr += "(?, ?, ?, ?),"
		vals = append(vals, car.CarID, car.CarModel, car.CarMake, car.CarModelYear)

	}

	sqlStr = sqlStr[0 : len(sqlStr)-1]
	stmt, _ := db.Prepare(sqlStr)
	_, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}
}

