package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-2024/controller/dto"
	"prova-2024/services"
)

type IngredientController struct {
	ingredientService services.IngredientService
}

func NewIngredientController(services services.IngredientService) *IngredientController {
	return &IngredientController{ingredientService: services}
}

func (controller *IngredientController) AddIngredient(ctx *gin.Context) {
	createIngredientRequest := dto.CreateIngredientsRequest{}
	err := ctx.ShouldBindJSON(&createIngredientRequest)
	if err != nil {
		panic(err)
	}

	controller.ingredientService.AddIngredient(createIngredientRequest.MapToModel())

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *IngredientController) GetAllByRecipeID(ctx *gin.Context) {

	value := controller.ingredientService.GetAllByRecipeID(ctx.Param("id"))

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   value,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *IngredientController) Delete(ctx *gin.Context) {
	controller.ingredientService.RemoveIngredient(ctx.Param("id"))

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
