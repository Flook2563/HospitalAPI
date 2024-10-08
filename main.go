package main

import (
	"log"

	"github.com/Flook2563/Hospitalapi/api"
	"github.com/Flook2563/Hospitalapi/config"
	"github.com/Flook2563/Hospitalapi/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDatabase()

	router := gin.Default()

	staffGroup := router.Group("/staff")
	{
		staffGroup.POST("/create", api.CreateStaff)
		staffGroup.POST("/login", api.LoginStaff)
	}

	patientGroup := router.Group("/patient")
	patientGroup.Use(middleware.JWTAuthMiddleware())
	{
		patientGroup.GET("/search", api.SearchPatient)
	}

	router.Run(":8080")

}
