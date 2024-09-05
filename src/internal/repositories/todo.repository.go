package repositories

import "0x5ea000000/ramos/internal/entities"

type TodoRepository interface {
	Create(todo *entities.Todo) (*entities.Todo, error)
	GetAll() ([]*entities.Todo, error)
	Get(id string) (*entities.Todo, error)
	Update(todo *entities.Todo) (*entities.Todo, error)
	Delete(id string) error
}
