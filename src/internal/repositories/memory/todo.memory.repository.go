package memory

import (
	"0x5ea000000/ramos/internal/entities"
	"errors"
	"gorm.io/gorm"
)

type TodoMemoryRepository struct {
	todos []*entities.Todo
}

// NewTodoRepository returns a new instance of TodoMemoryRepository
func NewTodoRepository(db *gorm.DB) *TodoMemoryRepository {
	return &TodoMemoryRepository{todos: []*entities.Todo{}}
}

func (r *TodoMemoryRepository) Create(todo *entities.Todo) (*entities.Todo, error) {
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *TodoMemoryRepository) GetAll() ([]*entities.Todo, error) {
	return r.todos, nil
}

func (r *TodoMemoryRepository) Get(id string) (*entities.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return &entities.Todo{}, errors.New("rating not found")
}

func (r *TodoMemoryRepository) Update(todo *entities.Todo) (*entities.Todo, error) {
	for idx, item := range r.todos {
		if item.ID == todo.ID {
			r.todos[idx].Name = todo.Name
			r.todos[idx].Note = todo.Note
			return r.todos[idx], nil
		}
	}
	return &entities.Todo{}, errors.New("rating not found")
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
