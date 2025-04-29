package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config almacena toda la configuración de la aplicación
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig almacena la configuración del servidor
type ServerConfig struct {
	Port int
	Mode string
}

// DatabaseConfig almacena la configuración de la base de datos
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetDSN devuelve el DSN para conectarse a la base de datos PostgreSQL
func (db DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode)
}

// LoadConfig carga la configuración desde variables de entorno
func LoadConfig() *Config {
	// Cargar variables desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env:", err)
	}

	// Valores por defecto
	config := &Config{
		Server: ServerConfig{
			Port: 8080,
			Mode: "debug",
		},
		Database: DatabaseConfig{
			Host:     "",
			Port:     5432,
			User:     "",
			Password: "",
			DBName:   "",
			SSLMode:  "",
		},
	}

	// Sobrescribir con variables de entorno si existen
	if port, exists := os.LookupEnv("SERVER_PORT"); exists {
		if p, err := strconv.Atoi(port); err == nil {
			config.Server.Port = p
		}
	}

	if mode, exists := os.LookupEnv("GIN_MODE"); exists {
		config.Server.Mode = mode
	}

	if host, exists := os.LookupEnv("DB_HOST"); exists {
		config.Database.Host = host
	}

	if port, exists := os.LookupEnv("DB_PORT"); exists {
		if p, err := strconv.Atoi(port); err == nil {
			config.Database.Port = p
		}
	}

	if user, exists := os.LookupEnv("DB_USER"); exists {
		config.Database.User = user
	}

	if password, exists := os.LookupEnv("DB_PASSWORD"); exists {
		config.Database.Password = password
	}

	if dbName, exists := os.LookupEnv("DB_NAME"); exists {
		config.Database.DBName = dbName
	}

	if sslMode, exists := os.LookupEnv("DB_SSL_MODE"); exists {
		config.Database.SSLMode = sslMode
	}

	// Imprimir la configuración para depuración (sin la contraseña)
	log.Printf("Configuración cargada: Server [Port: %d, Mode: %s], Database [Host: %s, Port: %d, User: %s, DBName: %s, SSLMode: %s]",
		config.Server.Port, config.Server.Mode,
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.DBName, config.Database.SSLMode)

	return config
}
