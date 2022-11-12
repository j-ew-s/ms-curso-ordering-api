package models

type Item struct {
	ProductId int32   `json:"product_id,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	Price     float64 `json:"price,omitempty"`
	TotalCost float64 `json:"total_cost,omitempty"`
	OrderId   int32   `json:"order_id,omitempty"`
}

type Order struct {
	Id                 int32   `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	UserId             int     `json:"user_id,omitempty"`
	Items              []Item  `json:"items,omitempty"`
	DiscountPercentage float64 `json:"discount_percentage,omitempty"`
	TotalCost          float64 `json:"total_cost,omitempty"`
}

type Orders struct {
	Orders []Order
}
