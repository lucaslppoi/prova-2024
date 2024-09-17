package repository

import (
	"gorm.io/gorm"
	"prova-2024/model"
)

type IngredientRepositoryImpl struct {
	Db *gorm.DB
}

func NewIngredientRepositoryImpl(Db *gorm.DB) IngredientRepository {
	return &IngredientRepositoryImpl{Db: Db}
}

func (r IngredientRepositoryImpl) AddIngredient(recipe model.Ingredient) {
	result := r.Db.Create(&recipe)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (r IngredientRepositoryImpl) GetAllByRecipeID(id string) []model.Ingredient {
	var values []model.Ingredient
	r.Db.Where("recipe_id = ?", id).Find(&values)
	return values
}

func (r IngredientRepositoryImpl) RemoveIngredient(id string) {
	result := r.Db.Delete(&model.Ingredient{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
}
