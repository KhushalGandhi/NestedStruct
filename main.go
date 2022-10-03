package main

import (
	"calling/models"
	"calling/routers"
	"fmt"
)

func main() {
	r := routers.RegisterRoutes()

	models.ConnecttoDatabase()

	fmt.Println("Successfully connected")
	r.Run("localhost:8080")
}
