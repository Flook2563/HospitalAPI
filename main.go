package main

import (
	"github.com/Flook2563/Hospitalapi/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	router.Run(":8080")

}
