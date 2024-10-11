package internal

import (
	"0x5ea000000/ramos/auth/internal/controllers"
	"0x5ea000000/ramos/auth/internal/services"
	"0x5ea000000/ramos/pkg/middlewares"
	"0x5ea000000/ramos/pkg/models"
	"0x5ea000000/ramos/pkg/repositories/database"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func Bootstrap() {
	r := gin.Default()

	r.Use(middlewares.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	dsn := "username:password@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Initialize database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&models.Todo{})

	repo := database.NewAuthRepository(db)
	service := services.NewAuthService(repo)
	controller := controllers.NewAuthController(service)

	// Create a new Todo
	r.POST("/register", controller.Register)

	// Get all Todos
	r.GET("/login", controller.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
