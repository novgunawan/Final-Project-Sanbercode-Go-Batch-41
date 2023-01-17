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
	router.POST("/customer", controllers.InsertCustomer)
	// POST
	// router.POST("/customer/verify", controllers.VerifyCustomer)
	// router.POST("/restaurant/verify", controllers.VerifyRestaurant)

	// router.POST("/customer", controllers.CreateCustomer)
	router.POST("/restaurant", controllers.InsertRestaurant)
	router.POST("/restaurant/insert-food", controllers.InsertFood)

	return router
}
