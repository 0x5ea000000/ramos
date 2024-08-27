package controllers

import (
	"0x5ea000000/ramos/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var newTodo entities.Todo

	// Bind JSON to the Todo struct
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new Todo to the todoList
	entities.TodoList = append(entities.TodoList, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// Get all Todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, entities.TodoList)
}

// Update a Todo
func UpdateTodo(c *gin.Context) {
	var updatedTodo entities.Todo
	name := c.Param("name")

	// Bind JSON to the Todo struct
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find and update the Todo
	for i, todo := range entities.TodoList {
		if todo.Name == name {
			entities.TodoList[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// Delete a Todo
func DeleteTodo(c *gin.Context) {
	name := c.Param("name")

	// Find and delete the Todo
	for i, todo := range entities.TodoList {
		if todo.Name == name {
			entities.TodoList = append(entities.TodoList[:i], entities.TodoList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
