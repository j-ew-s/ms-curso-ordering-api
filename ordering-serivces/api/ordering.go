package api

import (
	"github.com/j-ew-s/ms-curso-ordering-api/database"
	"github.com/valyala/fasthttp"
)

type Ordering struct {
	DBConn *database.SQLCommand
}

func (o Ordering) PlaceOrder(ctx *fasthttp.RequestCtx) {

}
