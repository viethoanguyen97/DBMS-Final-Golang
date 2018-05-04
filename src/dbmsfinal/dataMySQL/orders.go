package queryMySQL

type Order struct {
	OrderID    int64 `csv:"order_id" json:"order_id" omit:empty`
	CustomerID int64 `csv:"customer_id" json:"customer_id"`
}
