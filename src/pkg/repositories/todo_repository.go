package repositories

import (
	"0x5ea000000/ramos/pkg/models"
)

type TodoRepository interface {
	Create(todo *models.Todo) (*models.Todo, error)
	GetAll() ([]*models.Todo, error)
	Get(id string) (*models.Todo, error)
	Update(todo *models.Todo) (*models.Todo, error)
	Delete(id string) error
}
