package routers

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func StartServer() *gin.Engine {
	router := gin.Default()
	authorizedCustomer := router.Group("/", gin.BasicAuth(gin.Accounts{
		"cust": "cust",
	}))
	authorizedResto := router.Group("/", gin.BasicAuth(gin.Accounts{
		"resto": "resto",
	}))
	// GET
	authorizedResto.GET("/restaurant/foods", controllers.GetAllFood)
	authorizedCustomer.GET("/customer/foods", controllers.GetAllFood)

	authorizedCustomer.GET("/customer/verify/:phoneNumber", controllers.VerifyCustomer)
	authorizedCustomer.GET("/customer/:id/order/:status", controllers.GetActiveOrderResto)

	authorizedResto.GET("/restaurant/verify/:phoneNumber", controllers.VerifyRestaurant)
	authorizedResto.GET("/restaurant/:id/order/:status", controllers.GetActiveOrderResto)

	// POST

	authorizedCustomer.POST("/customer", controllers.InsertCustomer)
	authorizedCustomer.POST("/customer/insert-order", controllers.InsertOrder)
	authorizedResto.POST("/restaurant", controllers.InsertRestaurant)
	authorizedResto.POST("/restaurant/insert-food", controllers.InsertFood)
	authorizedResto.POST("/restaurant/:id/finish-order", controllers.FinishOrder)

	return router
}
