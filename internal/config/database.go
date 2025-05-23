package config

import (
	"fmt"
	"log"
	"time"

	"github.com/klmunoz2/grusolpac_api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB inicializa la conexión a la base de datos
func InitDB(config *Config) *gorm.DB {
	dsn := config.Database.GetDSN()

	// Log del DSN sin mostrar la contraseña
	//dsnForLog := config.Database.GetDSN()
	//log.Printf("Intentando conectar a la base de datos: %s", dsnForLog)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Intentar conectar con reintentos
	var db *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
		if err == nil {
			break
		}

		log.Printf("Intento %d de %d: Error al conectar a la base de datos: %v", i+1, maxRetries, err)
		if i < maxRetries-1 {
			log.Printf("Reintentando en 3 segundos...")
			time.Sleep(3 * time.Second)
		}
	}

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos después de %d intentos: %v", maxRetries, err)
	}

	// Verificar la conexión
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión SQL subyacente: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida exitosamente")
	return db

}

func AutoMigrate(db *gorm.DB) error {
	// Crear extensión pgcrypto si no existe
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"pgcrypto\"").Error; err != nil {
		return fmt.Errorf("failed to create pgcrypto extension: %w", err)
	}

	// Migrar modelos en el orden correcto
	modelsToMigrate := []interface{}{
		&models.Role{},
		&models.User{},
		&models.Student{},
		&models.Course{},
		&models.Module{},
		&models.Topic{},
		&models.Question{},
		&models.Answer{},
		&models.CourseEvaluation{},
		&models.EvaluationQuestion{},
		&models.Enrollment{},
		&models.Certificate{},
	}

	// Desactivar restricciones temporales para evitar problemas con el orden
	if err := db.Exec("SET CONSTRAINTS ALL DEFERRED").Error; err != nil {
		return fmt.Errorf("failed to defer constraints: %w", err)
	}

	// Ejecutar migración en transacción
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, model := range modelsToMigrate {
			if err := tx.AutoMigrate(model); err != nil {
				return fmt.Errorf("failed to migrate %T: %w", model, err)
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}

	log.Println("✅ Migración de base de datos completada exitosamente")
	return nil
}
