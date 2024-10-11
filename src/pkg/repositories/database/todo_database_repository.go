package database

import (
	"0x5ea000000/ramos/pkg/models"
	"gorm.io/gorm"
)

type TodoDatabaseRepository struct {
	*DatabaseRepository[models.Todo]
}

// NewTodoRepository returns a new instance of TodoDatabaseRepository
func NewTodoRepository(db *gorm.DB) *TodoDatabaseRepository {
	return &TodoDatabaseRepository{
		DatabaseRepository: NewDatabaseRepository[models.Todo](db),
	}
}
