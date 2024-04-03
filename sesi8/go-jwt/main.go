package main

import (
	"go-jwt/database"
	"go-jwt/router"
)

var (
	PORT = ":9090"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(PORT)
}
