package controller

import (
	"golang-saya/Models"
	stockmodivikasi "golang-saya/Models/StockModivikasi"
	stockmodivikasirepository "golang-saya/Repository/StokModivikasiRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStokModivikasiByID(c *gin.Context) {
	var request stockmodivikasi.StokModivikasi
	var response Models.BaseResponseModels

	response = stockmodivikasirepository.GetStokModivikasiByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertStokModivikasi(c *gin.Context) {
	var request stockmodivikasi.StokModivikasi
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

	response = stockmodivikasirepository.InsertStokModifikasi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateStokModivikasi(c *gin.Context) {
	var request stockmodivikasi.StokModivikasi
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

	response = stockmodivikasirepository.UpdateStokModivikasi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteStokModivikasi(c *gin.Context) {
	var request stockmodivikasi.StokModivikasi
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

	response = stockmodivikasirepository.DeleteStokModivikasi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
