package models

type Item struct {
	ProductId int32   `json:"product_id,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty"`
	TotalCost float64 `json:"total_cost,omitempty"`
	OrderId   int32   `json:"order_id,omitempty" `
}
