package orderingmain

import (
	api "github.com/j-ew-s/ms-curso-ordering-api/ordering-services/api"

	"github.com/valyala/fasthttprouter"
)

type OrderingMain struct {
	api *api
}

func SetOrderingMain() OrderingMain {

	api := &api{}
	entity := OrderingMain{
		api: api,
	}

	return entity
}

func (o OrderingMain) SetRoutes(router *fasthttprouter.Router) {
	router.POST("/", api.PlaceOrder)
}
