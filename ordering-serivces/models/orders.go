package models

type Item struct {
	ProductId int32
	Quantity  int
	Price     float64
	Discounts int
	TotalCost float64
}

type Orders struct {
	Id        int32
	UserId    int
	Items     []Item
	Discounts int
	TotalCost float64
}
