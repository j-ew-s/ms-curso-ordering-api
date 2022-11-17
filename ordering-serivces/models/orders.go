package models

type Order struct {
	Id                 int32   `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	UserId             int     `json:"user_id,omitempty"`
	Items              []Item  `json:"items,omitempty" gorm:"foreignKey:OrderId"`
	DiscountPercentage float64 `json:"discount_percentage,omitempty"`
	TotalCost          float64 `json:"total_cost,omitempty"`
}

type Orders struct {
	Orders []Order
}
