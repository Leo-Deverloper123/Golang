package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hospital-middleware/internal/handlers"
	"hospital-middleware/internal/middleware"
	"hospital-middleware/internal/models"
	"hospital-middleware/internal/services"
	"log"
	"os"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schemas
	db.AutoMigrate(&models.Hospital{}, &models.Staff{}, &models.Patient{})

	// Initialize services
	authService := services.NewAuthService(os.Getenv("JWT_SECRET"))
	staffService := services.NewStaffService(db)
	patientService := services.NewPatientService(db)

	// Initialize handlers
	staffHandler := handlers.NewStaffHandler(staffService, authService)
	patientHandler := handlers.NewPatientHandler(patientService)

	// Setup router
	router := gin.Default()

	// Public routes
	router.POST("/staff/create", staffHandler.Create)
	router.POST("/staff/login", staffHandler.Login)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(authService))
	{
		protected.GET("/patient/search", patientHandler.Search)
	}

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}