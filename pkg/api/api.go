package api

import (
	"github.com/gin-gonic/gin"
	contactHandler "github.com/kilianp07/CassandraCRUD/pkg/handlers/contact"
)

func Start() {

	// Run gin server
	router := gin.Default()

	router.GET("/contact", contactHandler.GetAll)
	router.GET("/contact/:id", contactHandler.GetById)
	router.POST("/contact", contactHandler.Create)
	router.PUT("/contact/:id", contactHandler.Update)
	router.DELETE("/contact/:id", contactHandler.Delete)

	router.Run(":8080")
}
