package internal

import (
	"0x5ea000000/ramos/pkg/models"
	"0x5ea000000/ramos/pkg/repositories/database"
	"0x5ea000000/ramos/todo/internal/controllers"
	"0x5ea000000/ramos/todo/internal/middlewares"
	"0x5ea000000/ramos/todo/internal/services"
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

	repo := database.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	controller := controllers.NewTodoController(service)

	// Create a new Todo
	r.POST("/todo", controller.Create)

	// Get all Todos
	r.GET("/todo", controller.Get)

	// Update a Todo by name
	r.PUT("/todo/:name", controller.Update)

	// Delete a Todo by name
	r.DELETE("/todo/:name", controller.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
