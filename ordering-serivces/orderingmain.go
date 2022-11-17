package orderingmain

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-ordering-api/database"
	"github.com/j-ew-s/ms-curso-ordering-api/ordering-serivces/api"
)

type OrderingMain struct {
	orderingAPI *api.Ordering
}

func CreateOrderingMain(dbConn *database.SQLCommand) OrderingMain {

	o := &api.Ordering{
		DBConn: dbConn,
	}

	orderingMain := OrderingMain{
		orderingAPI: o,
	}

	return orderingMain

}

func (o OrderingMain) SetRoutes(router *fasthttprouter.Router) {
	router.POST("/", o.orderingAPI.PlaceOrder)
	router.GET("/user/:id", o.orderingAPI.GetOrdersByUserId)
	router.GET("/detail/:id", o.orderingAPI.GetOrderById)
}
