package router

import (
	"github.com/gin-gonic/gin"
	"prova-2024/controller"
)

func NewRouter(recipeController *controller.RecipeController,
	ingredientController *controller.IngredientController) *gin.Engine {
	service := gin.Default()

	service.POST("/recipe", recipeController.AddRecipe)

	service.PUT("/recipe", recipeController.UpdateRecipe)
	service.GET("/recipe/id/:id", recipeController.GetByID)

	return service
}
