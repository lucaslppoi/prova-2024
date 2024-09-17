package services

import "prova-2024/model"

type RecipeService interface {
	Create(recipe model.Recipe) model.Recipe
	GetAll() []model.Recipe
	Update(recipe model.Recipe) model.Recipe
	Delete(id string)
}

type IngredientService interface {
	AddIngredient(item model.Ingredient)
	RemoveIngredient(id string)
	GetAllByRecipeID(id string) []model.Ingredient
}
