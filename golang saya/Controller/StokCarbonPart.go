package controller

import (
	"golang-saya/Models"
	stockcarbonpart "golang-saya/Models/StockCarbonPart"
	stokcarbonpartrepository "golang-saya/Repository/StokCarbonPartRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokCarbonPartByID(c *gin.Context) {
	var request stockcarbonpart.StokCarbonPark
	var response Models.BaseResponseModels

	response = stokcarbonpartrepository.GetStokCarbonPartByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertStokCarbonPart(c *gin.Context) {
	var request stockcarbonpart.StokCarbonPark
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

	response = stokcarbonpartrepository.InsertStokCarbonPart(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateStokCarbonPart(c *gin.Context) {
	var request stockcarbonpart.StokCarbonPark
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

	response = stokcarbonpartrepository.UpdateStokCarbonPart(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteStokCarbonPart(c *gin.Context) {
	var request stockcarbonpart.StokCarbonPark
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

	response = stokcarbonpartrepository.DeleteStokCarbonPart(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
