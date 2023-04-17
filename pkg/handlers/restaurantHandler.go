package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	restaurantRepo "github.com/kilianp07/CassandraCRUD/pkg/repositories/restaurant"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

// CreateRestaurant creates a new restaurant
func CreateRestaurant(c *gin.Context) {
	var restaurant structs.Restaurant
	if err := c.ShouldBindJSON(&restaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Implement logic to create restaurant in the database

	c.JSON(http.StatusOK, gin.H{"message": "Restaurant created successfully", "data": restaurant})
}

// GetRestaurant gets a restaurant by ID
func GetRestaurant(c *gin.Context) {
	var (
		restaurant   *structs.Restaurant
		err          error
		restaurantID int
	)

	// Cast ID parameter from string to int
	if restaurantID, err = strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Unable to process with this id parameter": err.Error()})
		return
	}

	// Get restaurant from cassandra DB
	if restaurant, err = restaurantRepo.GetRestaurantById(restaurantID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Cannot retrieve data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}
