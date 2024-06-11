package controller

import (
	"golang-saya/Models"
	KeteranganModels "golang-saya/Models/KeteranganModels"
	KeteranganBengkelRepository "golang-saya/Repository/KeteranganBengkelRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKeterangan(c *gin.Context) {
	var request KeteranganModels.KeteranganBengkel
	var response Models.BaseResponseModels

	response = KeteranganBengkelRepository.GetKeteranganByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertKeterangan(c *gin.Context) {
	var request KeteranganModels.KeteranganBengkel
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

	response = KeteranganBengkelRepository.InsertKeterangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateKeterangan(c *gin.Context) {
	var request KeteranganModels.KeteranganBengkel
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

	response = KeteranganBengkelRepository.UpdateKeterangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteKeterangan(c *gin.Context) {
	var request KeteranganModels.KeteranganBengkel
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

	response = KeteranganBengkelRepository.DeleteKeterangan(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
