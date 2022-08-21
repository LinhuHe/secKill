package datamodel

type Order struct {
	ID           int64 `sql:"id"`
	UserID       int64 `sql:"user_id"`
	ProductId    int64 `sql:"product_id"`
	OrderStataus int64 `sql:"order_status"`
}

const (
	OrderWait = iota
	OrderSuccess
	OrderFailed
)
