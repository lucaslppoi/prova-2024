package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-2024/controller/dto"
	"prova-2024/services"
)

type RecipeController struct {
	ingredientService services.RecipeService
}

func NewRecipeController(services services.RecipeService) *RecipeController {
	return &RecipeController{ingredientService: services}
}

func (controller *RecipeController) AddRecipe(ctx *gin.Context) {
	createRecipeRequest := dto.CreateRecipeRequest{}
	err := ctx.ShouldBindJSON(&createRecipeRequest)
	if err != nil {
		panic(err)
	}

	value := controller.ingredientService.Create(createRecipeRequest.MapToModel())

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   value,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RecipeController) UpdateRecipe(ctx *gin.Context) {
	createRecipeRequest := dto.CreateRecipeRequest{}
	err := ctx.ShouldBindJSON(&createRecipeRequest)
	if err != nil {
		panic(err)
	}

	value := controller.ingredientService.Update(createRecipeRequest.MapToModel())

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   value,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RecipeController) GetAll(ctx *gin.Context) {

	value := controller.ingredientService.GetAll()

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   value,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RecipeController) Delete(ctx *gin.Context) {
	controller.ingredientService.Delete(ctx.Param("id"))

	webResponse := dto.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
