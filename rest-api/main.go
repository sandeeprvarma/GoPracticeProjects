package main

// https://pkg.go.dev/

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.DbInit()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
