package api

import (
	"encoding/json"
	"fmt"

	"github.com/j-ew-s/ms-curso-ordering-api/common"
	"github.com/j-ew-s/ms-curso-ordering-api/database"
	"github.com/j-ew-s/ms-curso-ordering-api/ordering-serivces/models"
	"github.com/valyala/fasthttp"
)

type Ordering struct {
	DBConn *database.SQLCommand
}

func (o Ordering) PlaceOrder(ctx *fasthttp.RequestCtx) {

	order := models.Order{}

	err := json.Unmarshal(ctx.PostBody(), &order)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if order.DiscountPercentage >= 1 {
		common.PrepareResponse(ctx, 400, "Não pode aplicar 100% de descontos")
	}

	o.CalculateOrderTotal(&order)

	err = o.DBConn.InsertOrder(&order)

	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, 201, order.Id)

}

func (o Ordering) CalculateOrderTotal(order *models.Order) {

	for i, item := range order.Items {
		order.Items[i].TotalCost = (item.Price * float64(item.Quantity))
		order.TotalCost = order.TotalCost + order.Items[i].TotalCost
	}

	if order.DiscountPercentage > 0 {
		discountPrice := order.TotalCost * order.DiscountPercentage
		order.TotalCost = order.TotalCost - discountPrice
	}
}
