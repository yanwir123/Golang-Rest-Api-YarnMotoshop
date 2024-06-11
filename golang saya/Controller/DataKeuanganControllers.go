package controller

import (
	"golang-saya/Models"
	datakeuanganmodels "golang-saya/Models/DataKeuanganModels"
	DataKeuanganRepository "golang-saya/Repository/DataKeuanganRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataKeuanganByID(c *gin.Context) {
	var request datakeuanganmodels.DataKeuangan
	var response Models.BaseResponseModels

	response = DataKeuanganRepository.GetDataKeuanganByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertDataKeuangan(c *gin.Context) {
	var request datakeuanganmodels.DataKeuangan
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

	response = DataKeuanganRepository.InsertDataKeuangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateDataKeuangan(c *gin.Context) {
	var request datakeuanganmodels.DataKeuangan
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

	response = DataKeuanganRepository.UpdateDataKeuangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteDataKeuangan(c *gin.Context) {
	var request datakeuanganmodels.DataKeuangan
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

	response = DataKeuanganRepository.DeleteDataKeuangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
