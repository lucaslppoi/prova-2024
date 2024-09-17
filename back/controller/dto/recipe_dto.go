package dto

import "prova-2024/model"

type CreateRecipeRequest struct {
	Id              int     `json:"id"`
	Nome            string  `json:"nome"`
	TempoPreparo    int     `json:"tempo_preparo"`
	CustoAproximado float64 `json:"custo_aproximado"`
}

func (r *CreateRecipeRequest) MapToModel() model.Recipe {
	return model.Recipe{
		Id:              r.Id,
		Nome:            r.Nome,
		TempoPreparo:    r.TempoPreparo,
		CustoAproximado: r.CustoAproximado,
	}
}
