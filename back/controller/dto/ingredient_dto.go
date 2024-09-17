package dto

import "prova-2024/model"

type CreateIngredientsRequest struct {
	Id       int    `json:"id"`
	RecipeId int    `json:"recipe_id"`
	Nome     string `json:"nome"`
}

func (r *CreateIngredientsRequest) MapToModel() model.Ingredient {
	return model.Ingredient{
		Id:       r.Id,
		RecipeId: r.RecipeId,
		Nome:     r.Nome,
	}
}
