package models

type Item struct {
	ProductId   int     `json:"product_id,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	TotalCost   float64 `json:"total_cost,omitempty"`
}

type Order struct {
	UserId         int     `json:"user_id,omitempty"`
	Items          []Item  `json:"items,omitempty"`
	OrderTotalCost float64 `json:"order_total_cost,omitempty"`
}
