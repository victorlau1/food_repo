package main

import (
	"food-app/adapters"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	adapters.SetupRoutes(r)

	r.Run(":8080")
}
