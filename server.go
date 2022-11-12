package main

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-ordering-api/database"
	orderingmain "github.com/j-ew-s/ms-curso-ordering-api/ordering-serivces"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()
	sqlCommand := SetDataBase()

	orderingMain := orderingmain.CreateOrderingMain(sqlCommand)
	orderingMain.SetRoutes(router)

	fmt.Println("Ordering API : Up and running on port 5400.")
	fasthttp.ListenAndServe(":5400", CORS(router.Handler))

}

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS handles CORS
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func SetDataBase() *database.SQLCommand {

	conn := database.SetDrivers("mysql", "order_management")

	sqlCommand := &database.SQLCommand{
		SqlConn: conn,
	}

	err := sqlCommand.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return sqlCommand

}
