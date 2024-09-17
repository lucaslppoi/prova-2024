package repository

import (
	"gorm.io/gorm"
	"prova-2024/model"
)

type RecipeRepositoryImpl struct {
	Db *gorm.DB
}

func NewRecipeRepositoryImpl(Db *gorm.DB) RecipeRepository {
	return &RecipeRepositoryImpl{Db: Db}
}

func (r RecipeRepositoryImpl) Create(recipe model.Recipe) {
	result := r.Db.Create(&recipe)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (r RecipeRepositoryImpl) GetByID(id string) model.Recipe {
	var value model.Recipe
	r.Db.First(&value, id)
	return value
}

func (r RecipeRepositoryImpl) Update(recipe model.Recipe) {
	result := r.Db.Save(&recipe)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (r RecipeRepositoryImpl) Delete(id string) {
	result := r.Db.Delete(&model.Recipe{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
}
