package routers

import (
	"final-project/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func StartServer() *gin.Engine {
	router := gin.Default()

	// GET
	router.GET("/restaurant/foods", controllers.GetAllFood)
	router.GET("/customer/foods", controllers.GetAllFood)

	router.GET("/customer/verify/:phoneNumber", controllers.VerifyCustomer)
	router.GET("/restaurant/verify/:phoneNumber", controllers.VerifyRestaurant)

	// POST

	router.POST("/customer", controllers.InsertCustomer)
	router.POST("/restaurant", controllers.InsertRestaurant)
	router.POST("/restaurant/insert-food", controllers.InsertFood)

	return router
}
