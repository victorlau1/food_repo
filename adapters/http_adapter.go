package adapters

import (
	"food-app/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/code", RetrieveRecipes)
}

func RetrieveRecipes(c *gin.Context) {
	r := usecases.RetrieveRecipes()
	c.JSON(200, r)
}

func CreateRecipes(c *gin.Context) {
	var recipe usecases.Recipe

	if err := c.ShouldBind(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uc := usecases.CreateRecipe(recipe)
	c.JSON(201, uc)
	return
}
