package main

import (
	"deliveries/api"
	"deliveries/database"
)

func main() {
	database.Start()
	api.Start()
}
