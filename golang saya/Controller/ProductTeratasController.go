package controller

import (
	"golang-saya/Models"
	productteratasmodels "golang-saya/Models/ProductTeratasModels"
	productteratasrepository "golang-saya/Repository/ProductTeratasRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductTeratasByID(c *gin.Context) {
	var request productteratasmodels.ProductTeratas
	var response Models.BaseResponseModels

	response = productteratasrepository.GetProductTeratasByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertProductTeratas(c *gin.Context) {
	var request productteratasmodels.ProductTeratas
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

	response = productteratasrepository.InsertProductTeratas(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateProductTeratas(c *gin.Context) {
	var request productteratasmodels.ProductTeratas
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

	response = productteratasrepository.UpdateProductTeratas(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteProductTeratas(c *gin.Context) {
	var request productteratasmodels.ProductTeratas
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

	response = productteratasrepository.DeleteProductTeratas(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
