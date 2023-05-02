package api

import (
	"github.com/gin-gonic/gin"
	restaurantHandler "github.com/kilianp07/CassandraCRUD/pkg/handlers/restaurant"
)

func Start() {

	// Run gin server
	router := gin.Default()

	// Restaurant CRUD
	router.POST("/restaurant", restaurantHandler.CreateRestaurant)
	router.GET("/restaurant/:id", restaurantHandler.GetRestaurant)
	router.GET("/restaurant/all", restaurantHandler.GetAllRestaurants)
	router.DELETE("/restaurant/:id", restaurantHandler.DeleteRestaurant)
	router.PUT("/restaurant/:id", restaurantHandler.UpdateRestaurant)

	router.Run(":8080")
}
