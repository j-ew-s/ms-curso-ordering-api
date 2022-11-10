package main

import (
	"fmt"

	orderingmain "github.com/j-ew-s/ms-curso-ordering-api/ordering-services"
)

func main() {
	fmt.Print("ENTERED")
	orderingmain.Call()
}
