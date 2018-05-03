package queryMySQL

type Orderdetail struct {
	OrderID       int64 `csv:"order_id" json:"order_id"`
	CustomerID    int64 `json:"customer_id"`
	CarID         int64 `csv:"car_id" json:"car_id"`
	QuantityOrder int64 `csv:"quantity_order" json:"quantity_order"`
}

type OrderCarDetail struct {
	OrderID       int64  `json:"order_id"`
	CustomerID    int64  `json:"customer_id"`
	CarID         int64  `json:"car_id"`
	CarModel      string `json:"car_model"`
	QuantityOrder int64  `json:"quantity_order"`
}
