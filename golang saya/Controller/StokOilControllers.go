package controller

import (
	"golang-saya/Models"
	stockoilmodels "golang-saya/Models/StockOilModels"
	stockoilrepository "golang-saya/Repository/StockOilRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokOilByID(c *gin.Context) {
	var request stockoilmodels.StockOil
	var response Models.BaseResponseModels

	response = stockoilrepository.GetStokOilByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertStockOil(c *gin.Context) {
	var request stockoilmodels.StockOil
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = stockoilrepository.InsertStockOil(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateStockOil(c *gin.Context) {
	var request stockoilmodels.StockOil
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = stockoilrepository.UpdateStockOil(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteStokOil(c *gin.Context) {
	var request stockoilmodels.StockOil
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = stockoilrepository.DeleteStokOil(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
