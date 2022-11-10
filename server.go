package main

import (
	fasthttp "github.com/valyala/fasthttp"
	fasthttprouter "github.com/valyala/fasthttprouter"
)

func main() {

	router := fasthttprouter.New()

	// sqlCommand := SetDataBase()
	// clientLogin := loginServices.LoginserviceMain(sqlCommand)
	// loginServices.SetRoutes(router, clientLogin)

	// clientUser := userServices.UserServiceMain(sqlCommand)
	// userServices.SetRoutes(router, clientUser, clientLogin)

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

/*
func SetDataBase() *database.SQLCommand {

	conn := database.SetDrivers("mysql", "user_management")

	sqlCommand := &database.SQLCommand{
		SqlConn: conn,
	}

	err := sqlCommand.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return sqlCommand

}
*/
