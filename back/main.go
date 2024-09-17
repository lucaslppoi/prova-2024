package main

import (
	"net/http"
	"prova-2024/config"
	"prova-2024/controller"
	"prova-2024/model"
	"prova-2024/repository"
	"prova-2024/router"
	"prova-2024/services"
	"time"
)

func main() {
	db := config.DatabaseConnection()

	db.Table("recipes").AutoMigrate(&model.Recipe{})
	db.Table("ingredient").AutoMigrate(&model.Ingredient{})

	recipeRepository := repository.NewRecipeRepositoryImpl(db)
	ingredientsRepository := repository.NewIngredientRepositoryImpl(db)

	recipeService := services.NewRecipeServiceImpl(recipeRepository)
	ingredientsService := services.NewIngredientServiceImpl(ingredientsRepository)

	recipeController := controller.NewRecipeController(recipeService)
	ingredientController := controller.NewIngredientController(ingredientsService)

	routes := router.NewRouter(recipeController, ingredientController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
