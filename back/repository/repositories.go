package repository

import "prova-2024/model"

type IngredientRepository interface {
	AddIngredient(recipe model.Ingredient)
	GetByID(id string) model.Ingredient
	Update(recipe model.Ingredient)
	RemoveIngredient(id string)
}

type RecipeRepository interface {
	Create(recipe model.Recipe)
	GetByID(id string) model.Recipe
	Update(recipe model.Recipe)
	Delete(id string)
}
