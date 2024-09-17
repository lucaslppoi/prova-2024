package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"prova-2024/controller"
)

func NewRouter(recipeController *controller.RecipeController,
	ingredientController *controller.IngredientController) *gin.Engine {
	service := gin.Default()
	service.Use(cors.Default())

	service.POST("/receita", recipeController.AddRecipe)
	service.PUT("/receita", recipeController.UpdateRecipe)
	service.GET("/receita", recipeController.GetAll)
	service.DELETE("/receita/id/:id", recipeController.Delete)

	service.GET("/ingredients/id/:id", ingredientController.GetAllByRecipeID)
	service.POST("/ingredients", ingredientController.AddIngredient)
	service.DELETE("/ingredients/id/:id", ingredientController.Delete)

	return service
}
