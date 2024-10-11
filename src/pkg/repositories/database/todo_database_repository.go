package database

import (
	"0x5ea000000/ramos/pkg/models"
	"gorm.io/gorm"
)

type TodoDatabaseRepository struct {
	db *gorm.DB
}

// NewTodoRepository returns a new instance of TodoDatabaseRepository
func NewTodoRepository(db *gorm.DB) *TodoDatabaseRepository {
	return &TodoDatabaseRepository{db}
}

func (r *TodoDatabaseRepository) Create(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Create(todo).Error

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoDatabaseRepository) GetAll() ([]*models.Todo, error) {
	var todos []*models.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoDatabaseRepository) Get(id string) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoDatabaseRepository) Update(todo *models.Todo) (*models.Todo, error) {
	err := r.db.Save(todo).Error

	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoDatabaseRepository) Delete(id string) error {
	return r.db.Delete(&models.Todo{}, id).Error
}
