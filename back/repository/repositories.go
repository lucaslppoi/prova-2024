package repository

import "prova-2024/model"

type IngredientRepository interface {
	AddIngredient(recipe model.Ingredient)
	GetAllByRecipeID(id string) []model.Ingredient
	RemoveIngredient(id string)
}

type RecipeRepository interface {
	Create(recipe model.Recipe) model.Recipe
	GetAll() []model.Recipe
	Update(recipe model.Recipe) model.Recipe
	Delete(id string)
}
