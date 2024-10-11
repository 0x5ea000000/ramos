package memory

import (
	"0x5ea000000/ramos/pkg/models"
	"errors"
	"gorm.io/gorm"
)

type TodoMemoryRepository struct {
	todos []*models.Todo
}

// NewTodoRepository returns a new instance of TodoMemoryRepository
func NewTodoRepository(db *gorm.DB) *TodoMemoryRepository {
	return &TodoMemoryRepository{todos: []*models.Todo{}}
}

func (r *TodoMemoryRepository) Create(todo *models.Todo) (*models.Todo, error) {
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *TodoMemoryRepository) GetAll() ([]*models.Todo, error) {
	return r.todos, nil
}

func (r *TodoMemoryRepository) Get(id string) (*models.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return &models.Todo{}, errors.New("rating not found")
}

func (r *TodoMemoryRepository) Update(todo *models.Todo) (*models.Todo, error) {
	for idx, item := range r.todos {
		if item.ID == todo.ID {
			r.todos[idx].Name = todo.Name
			r.todos[idx].Note = todo.Note
			return r.todos[idx], nil
		}
	}
	return &models.Todo{}, errors.New("rating not found")
}

func (r *TodoMemoryRepository) Delete(id string) error {
	for idx, item := range r.todos {
		if item.ID == id {
			r.todos = append(r.todos[:idx], r.todos[idx+1:]...)
			return nil
		}
	}
	return errors.New("rating not found")
}
