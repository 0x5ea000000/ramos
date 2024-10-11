package controllers

import (
	"0x5ea000000/ramos/pkg/models"
	"0x5ea000000/ramos/todo/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (h *TodoController) Get(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (h *TodoController) Create(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTodo, err := h.service.Create(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTodo)
}

func (h *TodoController) Update(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRating, err := h.service.Update(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedRating)
}

func (h *TodoController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
