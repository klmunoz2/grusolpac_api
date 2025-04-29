package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(router *gin.Engine) {
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
		// Rutas de usuarios (ejemplo)
		userRoutes := v1.Group("/users")
		{
			userRoutes.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Lista de usuarios"})
			})
			userRoutes.GET("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{"message": "Detalles del usuario", "id": id})
			})
			userRoutes.POST("/", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado"})
			})
			userRoutes.PUT("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado", "id": id})
			})
			userRoutes.DELETE("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado", "id": id})
			})
		}

		// Puedes agregar más grupos de rutas aquí
	}

	// Ruta para health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
