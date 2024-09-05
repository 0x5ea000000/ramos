package database

import (
	"0x5ea000000/ramos/internal/entities"
	"gorm.io/gorm"
)

type TodoDatabaseRepository struct {
	db *gorm.DB
}

// NewTodoRepository returns a new instance of TodoDatabaseRepository
func NewTodoRepository(db *gorm.DB) *TodoDatabaseRepository {
	return &TodoDatabaseRepository{db}
}

func (r *TodoDatabaseRepository) Create(todo *entities.Todo) (*entities.Todo, error) {
	err := r.db.Create(todo).Error

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoDatabaseRepository) GetAll() ([]*entities.Todo, error) {
	var todos []*entities.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoDatabaseRepository) Get(id string) (*entities.Todo, error) {
	var todo entities.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoDatabaseRepository) Update(todo *entities.Todo) (*entities.Todo, error) {
	err := r.db.Save(todo).Error

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoDatabaseRepository) Delete(id string) error {
	return r.db.Delete(&entities.Todo{}, id).Error
}
