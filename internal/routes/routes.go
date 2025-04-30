package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klmunoz2/grusolpac_api/internal/handlers"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	// Ruta principal
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "API Grusolpac en funcionamiento",
			"version": "1.0.0",
			"author":  "klmunoz2",
		})
	})

	// Grupo de rutas API v1
	v1 := router.Group("/api/v1")
	{

		userRoutes := v1.Group("/users")
		{
			userRoutes.GET("/", userHandler.GetAllUsers)
			userRoutes.GET("/:id", userHandler.GetUserByID)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.PUT("/:id", userHandler.UpdateUser)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}

	}

	// Ruta para health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
