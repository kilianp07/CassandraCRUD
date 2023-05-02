package api

import (
	"github.com/gin-gonic/gin"
	addressHandler "github.com/kilianp07/CassandraCRUD/pkg/handlers/address"
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

	// Address CRUD
	router.POST("/address", addressHandler.CreateAddress)
	router.GET("/address/all", addressHandler.GetAllAddresses)
	router.GET("/address/:id", addressHandler.GetAddressById)
	router.PUT("/address/:id", addressHandler.UpdateAddress)
	router.DELETE("/address/:id", addressHandler.DeleteAddress)

	router.Run(":8080")
}
