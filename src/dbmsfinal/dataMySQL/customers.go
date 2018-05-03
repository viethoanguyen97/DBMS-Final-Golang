package queryMySQL

type Customer struct {
	CustomerID      int64  `csv:"customer_id" json:"customer_id"`
	CustomerName    string `csv:"customer_name" json:"customer_name"`
	CustomerEmail   string `csv:"customer_email" json:"customer_email"`
	CustomerAddress string `csv:"customer_address" json:"customer_address"`
}
