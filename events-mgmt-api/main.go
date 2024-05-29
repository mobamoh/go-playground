package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mobamoh/go-playground/events-mgmt-api/internal/db"
	"github.com/mobamoh/go-playground/events-mgmt-api/routes"
)

func main() {

	db.InitDB()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run()
}
