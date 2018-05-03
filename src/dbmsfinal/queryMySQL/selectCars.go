package queryMySQL

import (
	"errors"
	"fmt"
	"time"
)

func (r *CarsDAO) GetCarInfo(car_id int64) (*Car, float64, error) {
	start := time.Now()

	carInfo := &Car{}

	row := db.QueryRow("SELECT car_id, car_model, car_make, car_model_year FROM Cars WHERE car_id = ?;", car_id)

	elapsed := time.Since(start).Seconds()

	err := row.Scan(&carInfo.CarID, &carInfo.CarModel, &carInfo.CarMake, &carInfo.CarModelYear)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get car info")
	}

	return carInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCarsInfo() ([]*Car, float64, error) {
	start := time.Now()

	cars := make([]*Car, 0)

	rows, err := db.Query("SELECT car_id, car_model, car_make, car_model_year FROM Cars")

	elapsed := time.Since(start).Seconds()

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get all cars info")
	}

	cnt := 0
	for rows.Next() {
		carInfo := &Car{}
		cnt++
		err := rows.Scan(&carInfo.CarID, &carInfo.CarModel, &carInfo.CarMake, &carInfo.CarModelYear)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get all cars info")
		}

		cars = append(cars, carInfo)
	}

	defer rows.Close()

	return cars, elapsed, nil
}
