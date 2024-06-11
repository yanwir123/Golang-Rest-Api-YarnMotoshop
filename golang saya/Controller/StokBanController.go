package controller

import (
	"golang-saya/Models"
	stockbanmodels "golang-saya/Models/StockBanModels"
	stockbanrepository "golang-saya/Repository/StockBanRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokBanByID(c *gin.Context) {
	var request stockbanmodels.StockBan
	var response Models.BaseResponseModels

	response = stockbanrepository.GetStokBanByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertStockBan(c *gin.Context) {
	var request stockbanmodels.StockBan
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

	response = stockbanrepository.InsertStockBan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateStockBan(c *gin.Context) {
	var request stockbanmodels.StockBan
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

	response = stockbanrepository.UpdateStockBan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteStokBan(c *gin.Context) {
	var request stockbanmodels.StockBan
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

	response = stockbanrepository.DeleteStokBan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
