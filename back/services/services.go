package services

import "prova-2024/model"

type RecipeService interface {
	Create(recipe model.Recipe)
	GetByID(id string) model.Recipe
	Update(recipe model.Recipe)
	Delete(id string)
}

type IngredientService interface {
	AddIngredient(item model.Ingredient)
	RemoveIngredient(id string)
	Update(ingredient model.Ingredient)
	GetByID(id string) model.Ingredient
}
