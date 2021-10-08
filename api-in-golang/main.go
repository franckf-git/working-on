package main

import (
	"lite-api-crud/controllers"
	router "lite-api-crud/routers"
)

func main() {
	api := router.App{}
	api.Initialize()
	api.Run()
	defer controllers.Db.Close()
}
