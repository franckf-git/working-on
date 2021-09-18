package main

import (
	router "lite-api-crud/routers"
)

func main() {
	api := router.App{}
	api.Run()
}
