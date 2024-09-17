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

func (controller *IngredientController) UpdateRecipe(ctx *gin.Context) {
	createIngredientRequest := dto.CreateIngredientsRequest{}
	err := ctx.ShouldBindJSON(&createIngredientRequest)
	if err != nil {
		panic(err)
	}

	controller.ingredientService.Update(createIngredientRequest.MapToModel())

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *IngredientController) GetByID(ctx *gin.Context) {

	value := controller.ingredientService.GetByID(ctx.Param("id"))

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   value,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *IngredientController) Delete(ctx *gin.Context) {
	createIngredientRequest := dto.CreateRecipeRequest{}
	err := ctx.ShouldBindJSON(&createIngredientRequest)
	if err != nil {
		panic(err)
	}

	controller.ingredientService.RemoveIngredient(ctx.Param("id"))

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
