package addressHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	addressRepo "github.com/kilianp07/CassandraCRUD/pkg/repositories/address"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func CreateAddress(c *gin.Context) {
	var addressRequest structs.AddressRequest

	if err := c.BindJSON(&addressRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newAddress := &structs.Address{
		Id:           uuid.New().String(),
		Building:     addressRequest.Building,
		Street:       addressRequest.Street,
		Zipcode:      addressRequest.Zipcode,
		RestaurantId: addressRequest.RestaurantId,
	}

	if err := addressRepo.Create(newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Address created successfully", "created": newAddress})
}
