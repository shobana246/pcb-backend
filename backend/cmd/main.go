package main

import (
	"fmt"

	"pcb-backend/backend/database"
	"pcb-backend/backend/routes"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := routes.SetupRouter()

	fmt.Println("server start to run.....")
	router.Run(":8080")
}
