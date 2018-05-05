package queryMySQL

import (
	"errors"
	"fmt"
	"time"

	dataMySQL "dbmsfinal/dataMySQL"
)

type CarsDAO struct{}

func (r *CarsDAO) GetCarInfo(car_id int64) (*dataMySQL.Car, float64, error) {
	carInfo := &dataMySQL.Car{}

	start := time.Now()
	row := DB.QueryRow("SELECT car_id, car_model, car_make, car_model_year FROM Cars WHERE car_id = ?;", car_id)
	elapsed := time.Since(start).Seconds()

	err := row.Scan(&carInfo.CarID, &carInfo.CarModel, &carInfo.CarMake, &carInfo.CarModelYear)

	if err != nil {
		fmt.Println(err.Error())
		return nil, elapsed, errors.New("Fail to get car info")
	}

	return carInfo, elapsed, nil
}

func (r *CarsDAO) GetAllCarsInfo() (int64, float64, error) { //([]*dataMySQL.Car, float64, error) {
	start := time.Now()

	//	cars := make([]*dataMySQL.Car, 0)

	//rows, err := DB.Query("SELECT car_id, car_model, car_make, car_model_year FROM Cars")
	row := DB.QueryRow("SELECT count(*) FROM Cars")

	elapsed := time.Since(start).Seconds()

	var cnt int64
	err := row.Scan(&cnt)

	if err != nil {
		fmt.Println(err.Error())
		return 0, elapsed, err
		//return nil, elapsed, errors.New("Fail to get all cars info")
	}

	/*cnt := 0
	for rows.Next() {
		carInfo := &dataMySQL.Car{}
		cnt++
		err := rows.Scan(&carInfo.CarID, &carInfo.CarModel, &carInfo.CarMake, &carInfo.CarModelYear)

		if err != nil {
			fmt.Println(err.Error())
			return nil, elapsed, errors.New("Fail to get all cars info")
		}

		cars = append(cars, carInfo)
	}

	defer rows.Close()
	*/
	return cnt, elapsed, nil
	//return cars, elapsed, nil
}
