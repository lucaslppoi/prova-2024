package services

import (
	"prova-2024/model"
	"prova-2024/repository"
)

type RecipeServiceImpl struct {
	recipeRepository repository.RecipeRepository
}

func NewRecipeServiceImpl(recipeRepository repository.RecipeRepository) RecipeService {
	return &RecipeServiceImpl{
		recipeRepository: recipeRepository,
	}
}

func (r RecipeServiceImpl) Create(recipe model.Recipe) model.Recipe {
	return r.recipeRepository.Create(recipe)
}

func (r RecipeServiceImpl) GetAll() []model.Recipe {
	return r.recipeRepository.GetAll()
}

func (r RecipeServiceImpl) Update(recipe model.Recipe) model.Recipe {
	return r.recipeRepository.Update(recipe)
}

func (r RecipeServiceImpl) Delete(id string) {
	r.recipeRepository.Delete(id)
}
