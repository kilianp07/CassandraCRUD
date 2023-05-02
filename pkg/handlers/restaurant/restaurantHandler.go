package restaurantHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	restaurantRepo "github.com/kilianp07/CassandraCRUD/pkg/repositories/restaurant"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func CreateRestaurant(c *gin.Context) {
	var restaurantRequest structs.RestaurantRequest
	if err := c.BindJSON(&restaurantRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newRestaurant := &structs.Restaurant{
		Id:      uuid.New().String(),
		Borough: restaurantRequest.Borough,
		Cuisine: restaurantRequest.Cuisine,
		Name:    restaurantRequest.Name,
	}

	if err := restaurantRepo.Create(newRestaurant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Restaurant created successfully", "restaurant_id": newRestaurant.Id})
}

// GetAllRestaurants gets all restaurants
func GetAllRestaurants(c *gin.Context) {
	var (
		restaurants []*structs.Restaurant
		err         error
	)

	// Get all restaurants from cassandra DB
	if restaurants, err = restaurantRepo.GetAll(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot retrieve data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurants})
}

// GetRestaurant gets a restaurant by ID
func GetRestaurant(c *gin.Context) {
	var (
		restaurant *structs.Restaurant
		err        error
	)

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
		return
	}

	// Get restaurant from cassandra DB
	if restaurant, err = restaurantRepo.GetById(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot retrieve data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": restaurant})
}

func DeleteRestaurant(c *gin.Context) {
	var (
		err error
	)

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
		return
	}

	// Delete restaurant from cassandra DB
	if err = restaurantRepo.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot delete data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Restaurant deleted successfully"})
}

func UpdateRestaurant(c *gin.Context) {
	var (
		restaurantRequest structs.Restaurant
		err               error
		updatedRestaurant *structs.Restaurant
	)

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
		return
	}

	if err = c.BindJSON(&restaurantRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid input"})
		return
	}

	// Update restaurant in cassandra DB
	if updatedRestaurant, err = restaurantRepo.Update(restaurantRequest, c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot update data: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Restaurant updated successfully",
		"data":    updatedRestaurant,
	})
}
