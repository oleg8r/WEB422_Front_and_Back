package controllers

import (
	"booksApi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Create(context *gin.Context) {
	var drug models.Drug

	err := context.ShouldBind(&drug)

	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = models.Db.Create(&drug).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
	}

	context.JSON(http.StatusOK, drug)
}

func Read(context *gin.Context) {
	var drugs []models.Drug

	err := models.Db.Find(&drugs).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, nil)
	}

	context.JSON(http.StatusOK, drugs)
}

func Update(context *gin.Context) {
	var drug models.Drug
	id := context.PostForm("id")
	quantity := context.PostForm("quantity")

	err := models.Db.First(&drug, id).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	drug.Quantity, err = strconv.Atoi(quantity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	models.Db.Save(&drug)

	context.JSON(http.StatusOK, drug)
}

func Delete(context *gin.Context) {
	id := context.Query("id")
	err := models.Db.Delete(&models.Drug{}, id).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, "")

}
