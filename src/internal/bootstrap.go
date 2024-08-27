package internal

import (
	"0x5ea000000/ramos/internal/controllers"
	"0x5ea000000/ramos/internal/middlewares"
	"github.com/gin-gonic/gin"
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

	// Create a new Todo
	r.POST("/todo", controllers.CreateTodo)

	// Get all Todos
	r.GET("/todo", controllers.GetTodos)

	// Update a Todo by name
	r.PUT("/todo/:name", controllers.UpdateTodo)

	// Delete a Todo by name
	r.DELETE("/todo/:name", controllers.DeleteTodo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
