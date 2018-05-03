package queryMySQL

type Car struct {
	CarID        int64  `csv:"car_id" json:"car_id"`
	CarModel     string `csv:"car_model" json:"car_model"`
	CarMake      string `csv:"car_make" json:"car_make"`
	CarModelYear int    `csv:"car_model_year" json:"car_model_year"`
}
