package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFood(ctx *gin.Context) {
	var result gin.H
	food, err := repository.GetAllFood(database.Connection)
	if err != nil {
		result = gin.H{
			"success": false,
			"error":   err,
			"result":  food,
		}
	} else {
		result = gin.H{
			"success": true,
			"result":  food,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func InsertFood(ctx *gin.Context) {
	var food model.Foods
	var result gin.H

	if err := ctx.BindJSON(&food); err != nil {
		return
	}

	err := repository.InsertFood(database.Connection, food)
	service.CheckError(err)

	result = gin.H{
		"success": true,
		"result":  food,
	}
	ctx.JSON(http.StatusCreated, result)
}
