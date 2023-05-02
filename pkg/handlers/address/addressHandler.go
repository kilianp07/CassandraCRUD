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

func GetAllAddresses(c *gin.Context) {
	addresses, err := addressRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Addresses retrieved successfully", "addresses": addresses})
}

func GetAddressById(c *gin.Context) {
	address, err := addressRepo.GetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Address retrieved successfully", "address": address})
}

func UpdateAddress(c *gin.Context) {
	var (
		addressRequest structs.AddressRequest
		address        *structs.Address
		err            error
	)

	if err := c.BindJSON(&addressRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	updatedAddress := &structs.Address{
		Id:           c.Param("id"),
		Building:     addressRequest.Building,
		Street:       addressRequest.Street,
		Zipcode:      addressRequest.Zipcode,
		RestaurantId: addressRequest.RestaurantId,
	}

	if address, err = addressRepo.Update(updatedAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Address updated successfully", "updated": address})
}

func DeleteAddress(c *gin.Context) {

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid ID"})
		return
	}

	if err := addressRepo.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "Address deleted successfully"})
}
