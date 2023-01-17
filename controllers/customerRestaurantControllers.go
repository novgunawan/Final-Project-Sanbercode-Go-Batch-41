package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyCustomer(ctx *gin.Context) {
	var result gin.H
	var customer model.Customers
	var phoneNumber = ctx.Param("phoneNumber")
	// if err := ctx.ShouldBindJSON(&body); err != nil {
	// 	result = gin.H{
	// 		"error":   err.Error(),
	// 		"message": "Invalid input",
	// 	}
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, result)
	// }

	// if !condition {
	// 	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
	// 		"error":   "data not found",
	// 		"message": fmt.Sprintf("customer with phone number %s is not found", body.PhoneNumber),
	// 	})
	// 	return
	// }
	err := repository.VerifyCustomer(database.Connection, phoneNumber, customer)
	// service.CheckError(err)
	if err == nil {
		result = gin.H{
			"success": true,
			"message": "Customer with number" + phoneNumber + " is found.",
			"result":  customer,
		}
		ctx.JSON(http.StatusCreated, result)

	} else {
		result = gin.H{
			"success": false,
			"error":   err,
		}
		ctx.JSON(http.StatusNotFound, result)
	}

}

func InsertCustomer(ctx *gin.Context) {
	var result gin.H
	var customer model.Customers
	if err := ctx.BindJSON(&customer); err != nil {
		return
	}

	err := repository.InsertCustomer(database.Connection, customer)
	service.CheckError(err)

	result = gin.H{
		"success": true,
		"result":  customer,
	}
	ctx.JSON(http.StatusCreated, result)
}

func InsertRestaurant(ctx *gin.Context) {
	var result gin.H
	var restaurant model.Restaurants
	if err := ctx.BindJSON(&restaurant); err != nil {
		return
	}

	err := repository.InsertRestaurant(database.Connection, restaurant)
	service.CheckError(err)

	result = gin.H{
		"success": true,
		"result":  restaurant,
	}
	ctx.JSON(http.StatusCreated, result)
}
