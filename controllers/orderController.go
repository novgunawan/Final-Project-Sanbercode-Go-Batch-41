package controllers

import (
	"final-project/database"
	"final-project/model"
	"final-project/repository"
	"final-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertOrder(ctx *gin.Context) {
	var result gin.H
	var order model.Orders

	if err := ctx.BindJSON(&order); err != nil {
		result = gin.H{
			"success": false,
			"error":   err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, result)
	}
	status, err := repository.CreateOrder(database.Connection, order)
	service.CheckError(err)
	totalPrice := calculateOrder(order.Food)

	jsonResult := map[string]any{}
	jsonResult["totalPrice"] = totalPrice
	jsonResult["status"] = status
	result = gin.H{
		"success": true,
		"result":  jsonResult,
	}
	ctx.JSON(http.StatusCreated, result)
}

func calculateOrder(food []model.FoodOrder) int {
	var total = 0
	for _, f := range food {
		total += f.Price * f.Quantity
	}
	return total
}

func GetActiveOrderResto(ctx *gin.Context) {
	var result gin.H
	var order model.Orders
	var stringRestoID = ctx.Param("id")
	restoID, _ := strconv.Atoi(stringRestoID)
	var status = ctx.Param("status")
	order, err := repository.GetActiveOrderResto(database.Connection, restoID, status, order)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Error in finding active order.",
		}
		ctx.JSON(http.StatusCreated, result)

	} else {
		result = gin.H{
			"success": true,
			"result":  order,
		}
		ctx.JSON(http.StatusNotFound, result)

	}

}

func GetActiveOrderCustomer(ctx *gin.Context) {
	var result gin.H
	var order model.Orders
	var stringCustomerID = ctx.Param("id")
	customerID, _ := strconv.Atoi(stringCustomerID)
	var status = ctx.Param("status")
	order, err := repository.GetActiveOrderCustomer(database.Connection, customerID, status, order)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "Error in finding active order.",
		}
		ctx.JSON(http.StatusNotFound, result)

	} else {
		result = gin.H{
			"success": true,
			"result":  order,
		}
		ctx.JSON(http.StatusOK, result)

	}

}

type FinishOrderBody struct {
	Status string `json:"status"`
}

func FinishOrder(ctx *gin.Context) {
	var result gin.H
	var order model.Orders
	bodyStatus := FinishOrderBody{}
	if err := ctx.ShouldBindJSON(&bodyStatus); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	count := repository.FinishOrder(database.Connection, order, &bodyStatus.Status, id)

	if count == 0 {
		result = gin.H{
			"success": false,
			"error":   "error finishing order.",
		}
		ctx.JSON(http.StatusBadRequest, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Order is paid by customer.",
		}
		ctx.JSON(http.StatusOK, result)

	}

}
