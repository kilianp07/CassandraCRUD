package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kilianp07/CassandraCRUD/pkg/handlers"
)

func Start() {

	// Run gin server
	router := gin.Default()

	// Restaurant CRUD
	router.POST("/restaurant", handlers.CreateRestaurant)
	router.GET("/restaurant/:id", handlers.GetRestaurant)
	router.GET("/restaurant/all", handlers.GetAllRestaurants)
	router.DELETE("/restaurant/:id", handlers.DeleteRestaurant)
	router.PUT("/restaurant/:id", handlers.UpdateRestaurant)
	router.Run(":8080")
}
