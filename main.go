package main

import (
	"fmt"

	"github.com/PiterWeb/CalendarWebServer/routes"
)

func main() {

	fmt.Println("Server Running on port:")

	routes.Routes()

}
