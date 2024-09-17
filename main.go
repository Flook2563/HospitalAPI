package main

import (
	"github.com/Flook2563/Hospitalapi/api"
	"github.com/Flook2563/Hospitalapi/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()

	staffGroup := router.Group("/staff")
	{
		staffGroup.POST("/create", api.CreateStaff)
	}

	router.Run(":8080")

}
