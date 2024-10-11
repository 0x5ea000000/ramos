package database

import (
	"0x5ea000000/ramos/pkg/models"
	"gorm.io/gorm"
)

type AuthDatabaseRepository struct {
	*DatabaseRepository[models.User]
}

// NewAuthRepository returns a new instance of AuthDatabaseRepository
func NewAuthRepository(db *gorm.DB) *AuthDatabaseRepository {
	return &AuthDatabaseRepository{
		DatabaseRepository: NewDatabaseRepository[models.User](db),
	}
}
