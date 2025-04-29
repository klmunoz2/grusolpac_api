package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/klmunoz2/grusolpac_api/internal/config"
	"github.com/klmunoz2/grusolpac_api/internal/routes"
)

func main() {
	// Mostrar el directorio de trabajo actual para depuración
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error al obtener el directorio de trabajo: %v", err)
	} else {
		log.Printf("Directorio de trabajo actual: %s", dir)

		// Verificar si existe el archivo .env
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); os.IsNotExist(err) {
			log.Printf("Advertencia: El archivo .env no existe en %s", envPath)
		} else {
			log.Printf("El archivo .env existe en %s", envPath)
		}
	}

	// Cargar configuración
	log.Println("Cargando configuración...")
	cfg := config.LoadConfig()

	
	// Configurar modo de Gin
	gin.SetMode(cfg.Server.Mode)
	log.Printf("Modo Gin establecido a: %s", cfg.Server.Mode)

	// Inicializar la conexión a la base de datos
	log.Println("Inicializando conexión a la base de datos...")
	db := config.InitDB(cfg)
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión SQL subyacente: %v", err)
	}
	defer sqlDB.Close()

	// Inicializar el router de Gin
	log.Println("Inicializando router...")
	router := gin.Default()

	// Configurar rutas
	log.Println("Configurando rutas...")
	routes.SetupRoutes(router)

	// Iniciar el servidor
	serverAddr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("Servidor iniciando en http://localhost%s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
