package queryMySQL

type Order struct {
	OrderID    int64 `csv:"order_id"`
	CustomerID int64 `csv:"customer_id"`
}

type OrderDAO struct{}
