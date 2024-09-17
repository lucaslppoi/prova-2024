package services

import (
	"prova-2024/model"
	"prova-2024/repository"
)

type IngredientServiceImpl struct {
	ingredientRepository repository.IngredientRepository
}

func NewIngredientServiceImpl(ingredientRepository repository.IngredientRepository) IngredientService {
	return &IngredientServiceImpl{
		ingredientRepository: ingredientRepository,
	}
}

func (r IngredientServiceImpl) AddIngredient(ingredient model.Ingredient) {
	r.ingredientRepository.AddIngredient(ingredient)
}

func (r IngredientServiceImpl) GetAllByRecipeID(id string) []model.Ingredient {
	return r.ingredientRepository.GetAllByRecipeID(id)
}

func (r IngredientServiceImpl) RemoveIngredient(id string) {
	r.ingredientRepository.RemoveIngredient(id)
}
