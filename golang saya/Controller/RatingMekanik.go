package controller

import (
	"golang-saya/Models"
	RatingMekanik "golang-saya/Models/RatingMekanik"
	ratingmekanikrepository "golang-saya/Repository/RatingMekanikRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRatingMekanikByID(c *gin.Context) {
	var request RatingMekanik.RatingMekanik
	var response Models.BaseResponseModels

	response = ratingmekanikrepository.GetRatingMekanikByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertRatingMekanik(c *gin.Context) {
	var request RatingMekanik.RatingMekanik
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

	response = ratingmekanikrepository.InsertRatingMekanik(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateRatingMekanik(c *gin.Context) {
	var request RatingMekanik.RatingMekanik
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

	response = ratingmekanikrepository.UpdateRatingMekanik(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteRatingMekanik(c *gin.Context) {
	var request RatingMekanik.RatingMekanik
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

	response = ratingmekanikrepository.DeleteRatingMekanik(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
