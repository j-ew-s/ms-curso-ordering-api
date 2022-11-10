package orderingmain

import (
	"github.com/j-ew-s/ms-curso-ordering-api/src/ordering-services/api"

	fasthttprouter "github.com/valyala/fasthttprouter"
)

type OrderingMain struct {
	Ordering *api.Ordering
}

func SetOrderingMain() OrderingMain {

	o := &api.Ordering{}
	entity := OrderingMain{
		Ordering: o,
	}

	return entity
}

func (o OrderingMain) SetRoutes(router *fasthttprouter.Router) {
	router.POST("/", api.PlaceOrder)
}
